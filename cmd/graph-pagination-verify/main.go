package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultGraphAPI  = "http://127.0.0.1:8080/graphql"
	baycCollectionID = "612ecc22-36ef-4ef7-bb0b-5b864b85d089"

	defaultPageSize = 250
)

func main() {
	err := run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func run() error {

	var (
		flagDatabase     string
		flagLogLevel     string
		flagGraphAPI     string
		flagCollectionID string
		flagPageSize     uint
		flagDumpFile     string
	)
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
	pflag.StringVarP(&flagDatabase, "database", "d", "", "database address")
	pflag.StringVarP(&flagGraphAPI, "graph-api", "g", defaultGraphAPI, "Graph API endpoint")
	pflag.StringVarP(&flagCollectionID, "collection-id", "c", baycCollectionID, "collection ID")
	pflag.UintVarP(&flagPageSize, "page-size", "p", defaultPageSize, "number of items to request in a single page")
	pflag.StringVar(&flagDumpFile, "dump-file", "", "dump file to write response data to")

	pflag.Parse()

	// Initialize logging.
	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	level, err := zerolog.ParseLevel(flagLogLevel)
	if err != nil {
		return fmt.Errorf("could not parse log level: %w", err)
	}
	log = log.Level(level)
	zerolog.SetGlobalLevel(level)

	// Connect to the DB.
	dbCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(postgres.Open(flagDatabase), dbCfg)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	// Open the dump file, if needed.
	var dumpFile io.Writer = io.Discard
	if flagDumpFile != "" {
		dumpFile, err = os.Create(flagDumpFile)
		if err != nil {
			return fmt.Errorf("could not open dump file: %w", err)
		}
	}

	// Get the NFT IDs.
	ids, err := getCollectionNFTs(db, flagCollectionID)
	if err != nil {
		return fmt.Errorf("could not retrieve collection IDs: %w", err)
	}

	collectionSize := len(ids)

	log.Info().
		Int("size", collectionSize).
		Str("collection", flagCollectionID).
		Msg("retrieved NFT IDs")

	cursor := ""
	processed := 0
	pages := 0

	for {

		log.Info().Int("page", pages+1).Msg("requesting page")

		// Retrieve the page of NFTs.
		nftConn, err := getPageFromAPI(flagGraphAPI, flagCollectionID, dumpFile, flagPageSize, cursor)
		if err != nil {
			return fmt.Errorf("could not retrieve NFTs from the API: %w", err)
		}

		nftsLeft := uint(len(ids) - processed)
		batchSize := uint(len(nftConn.Edges))

		log.Info().
			Str("cursor", cursor).
			Int("processed", processed).
			Uint("got", batchSize).
			Uint("left", nftsLeft).
			Msg("requesting NFT page")

		// Check if we have the expected number of NFTs.
		expected := nftsLeft
		if nftsLeft > flagPageSize {
			expected = flagPageSize
		}
		if batchSize != expected {
			log.Error().
				Uint("expected", expected).
				Uint("got", batchSize).
				Msg("got unexpected number of NFTs")
			break
		}

		// Process results.
		for i, nft := range nftConn.Edges {

			idx := processed + i

			// Verify that the NFT is the expected one.
			if nft.Node.ID != ids[idx] {
				log.Error().
					Int("index", idx).
					Str("expected", ids[idx]).
					Str("got", nft.Node.ID).
					Msg("unexpected NFT ID")
			}

			// Verify that the cursor is correct.
			expectedCursor := base64Encode(ids[idx])
			if nft.Cursor != expectedCursor {
				log.Error().
					Int("index", idx).
					Str("expected", expectedCursor).
					Str("got", nft.Cursor).
					Msg("unexpected cursor")
			}
		}

		// Prepare parameters for the next page.
		processed += len(nftConn.Edges)
		pages++
		cursor = nftConn.Edges[len(nftConn.Edges)-1].Cursor

		// See if there should be more more pages.
		more := processed < collectionSize
		if more != nftConn.PageInfo.HasNextPage {
			log.Error().Bool("expected", more).Bool("got", nftConn.PageInfo.HasNextPage).Msg("unexpected hasNextPage")
			break
		}
		if !more {
			log.Info().Msg("processed all tokens")
			break
		}
	}

	log.Info().Int("total_pages", pages).Msg("processed all results")

	return nil
}
