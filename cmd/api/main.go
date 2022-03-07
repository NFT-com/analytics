package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	gormzerolog "github.com/wei840222/gorm-zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/NFT-com/indexer-api/api"
	"github.com/NFT-com/indexer-api/graph/generated"
	"github.com/NFT-com/indexer-api/storage"
)

const (
	// FIXME: Specify the default database
	defaultDatabase       = "host=localhost user=nft-user password=nft-test-pass dbname=nft-com port=5432 sslmode=disable"
	defaultPlaygroundPath = "/"
)

const (
	success = 0
	failure = 1
)

func main() {
	os.Exit(run())
}

func run() int {

	var (
		flagBind       string
		flagPlayground string
		flagDatabase   string
		flagLogLevel   string
	)

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagPlayground, "playground-path", "p", defaultPlaygroundPath, "path for GraphQL playground")
	pflag.StringVarP(&flagDatabase, "database", "d", defaultDatabase, "database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")

	pflag.Parse()

	// Initialize logging.
	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	log := zerolog.New(os.Stderr).With().Timestamp().Str("level", flagLogLevel).Logger()
	level, err := zerolog.ParseLevel(flagLogLevel)
	if err != nil {
		log.Error().Err(err).Msg("could not parse log level")
		return failure
	}
	log = log.Level(level)

	// FIXME: Quick choice for Gorm + zerolog, not a definite one.
	dbCfg := gorm.Config{
		Logger: gormzerolog.NewWithLogger(log),
	}
	db, err := gorm.Open(postgres.Open(flagDatabase), &dbCfg)
	if err != nil {
		log.Error().Err(err).Msg("could not connect to database")
		return failure
	}

	storage := storage.New(db)
	server := api.NewServer(storage)
	cfg := generated.Config{
		Resolvers: server,
	}
	schema := generated.NewExecutableSchema(cfg)
	gqlServer := handler.NewDefaultServer(schema)

	// FIXME: Remove this in a final version
	http.Handle(flagPlayground, playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", gqlServer)

	playgroundURL := formatPlaygroundURL(flagBind, flagPlayground)
	log.Info().Str("address", playgroundURL).Msg("GraphQL playground URL")

	err = http.ListenAndServe(flagBind, nil)
	if err != nil {
		log.Error().Err(err).Msg("could not start server")
		return failure
	}

	return success
}

func formatPlaygroundURL(address string, path string) string {

	path = strings.TrimPrefix(path, "/")

	if strings.HasPrefix(address, ":") {
		return fmt.Sprintf("http://localhost%s/%s", address, path)
	}

	return fmt.Sprintf("http://%s/%s", address, path)
}
