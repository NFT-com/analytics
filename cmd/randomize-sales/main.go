package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/spf13/pflag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	var (
		flagGraphDatabase  string
		flagEventsDatabase string
		flagBatchSize      uint
		flagDBLogLevel     string
	)

	pflag.StringVarP(&flagEventsDatabase, "events-database", "e", "", "events database address")
	pflag.StringVarP(&flagGraphDatabase, "graph-database", "g", "", "graph database address")
	pflag.UintVarP(&flagBatchSize, "batch-size", "b", 100_000, "batch size for events update")
	pflag.StringVarP(&flagDBLogLevel, "db-log-level", "l", "silent", "log level for the database logger")

	pflag.Parse()

	// Set database logger.
	logLevel := parseLogLevel(flagDBLogLevel)

	dblog := logger.New(
		log.New(os.Stderr, "\r\n [gorm] ", log.LstdFlags),
		logger.Config{
			Colorful: false,
			LogLevel: logLevel,
		},
	)
	dbCfg := &gorm.Config{
		Logger:                 dblog,
		SkipDefaultTransaction: true, // We'll create our own transactions.
	}

	// Open the Graph database.
	graphDB, err := gorm.Open(postgres.Open(flagGraphDatabase), dbCfg)
	if err != nil {
		log.Fatalf("could not open graph database: %s", err)
	}

	// Open the Events database.
	edb, err := gorm.Open(postgres.Open(flagEventsDatabase), dbCfg)
	if err != nil {
		log.Fatalf("could not open graph database: %s", err)
	}

	// Get list of collection IDs.
	collections, err := getCollections(graphDB)
	if err != nil {
		log.Fatalf("could not retrieve list of collections: %s", err)
	}

	log.Printf("retrieved %v collections: %+#v", len(collections), collections)

	// Get token IDs.
	tokens, err := getTokenMap(graphDB, collections)
	if err != nil {
		log.Fatalf("could not retrieve list of tokens: %s", err)
	}

	for collection, tokenID := range tokens {
		fmt.Printf("collection: %v token count: %v\n", collection, len(tokenID))
	}

	// FIXME: Remove when ready
	var count uint64
	err = edb.Raw("SELECT COUNT(*) FROM sales WHERE collection_address = '' AND token_id = ''").Scan(&count).Error
	if err == nil {
		log.Printf("total %v unassigned sales", count)
	}

	// Update events
	for {

		var saleIDs []string
		err = edb.Raw("SELECT id FROM sales WHERE collection_address = '' AND token_id = '' LIMIT ?", flagBatchSize).Scan(&saleIDs).Error
		if err != nil {
			log.Fatalf("could not retrieve sale IDs: %s", err)
		}

		log.Printf("retrieved %v unassigned sales", len(saleIDs))

		var queries []string
		for _, sale := range saleIDs {

			// Pick a random collectionID
			collection := collections[rand.Intn(len(collections))].Address

			// Pick a random token ID.
			tokenIDs := tokens[collection]
			tokenID := tokenIDs[rand.Intn(len(tokenIDs))]

			query := fmt.Sprintf("UPDATE SALES SET collection_address = '%s', token_id = '%s' WHERE id = '%s'", collection, tokenID, sale)
			queries = append(queries, query)
		}

		err = edb.Transaction(func(tx *gorm.DB) error {

			for _, query := range queries {
				err = tx.Exec(query).Error
				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			log.Fatalf("could not update sales: %s", err)
		}

		log.Printf("assigned %v sales", len(saleIDs))

		// If we retrieved less events than requested, this is the last batch.
		if len(saleIDs) < int(flagBatchSize) {
			log.Printf("last batch processed")
			break
		}
	}
}

func parseLogLevel(level string) logger.LogLevel {
	switch level {
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	case "silent":
		return logger.Silent

	default:
		return logger.Silent
	}
}
