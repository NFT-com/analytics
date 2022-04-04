package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	gormzerolog "github.com/wei840222/gorm-zerolog"
	"github.com/ziflex/lecho/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/NFT-com/graph-api/graph-api/api"
	"github.com/NFT-com/graph-api/graph-api/graph/generated"
	"github.com/NFT-com/graph-api/graph-api/storage"
)

const (
	defaultPlaygroundPath  = "/"
	defaultGraphQLEndpoint = "/graphql"
)

const (
	success = 0
	failure = 1
)

func main() {
	os.Exit(run())
}

func run() int {

	// Signal catching for clean shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	var (
		flagBind               string
		flagDatabase           string
		flagLogLevel           string
		flagPlayground         string
		flagComplexityLimit    int
		flagEnablePlayground   bool
		flagEnableQueryLogging bool
	)

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagDatabase, "database", "d", "", "database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
	pflag.StringVarP(&flagPlayground, "playground-path", "p", defaultPlaygroundPath, "path for GraphQL playground")
	pflag.IntVar(&flagComplexityLimit, "query-complexity", 0, "GraphQL query complexity limit")
	pflag.BoolVar(&flagEnablePlayground, "enable-playground", false, "enable GraphQL playground")
	pflag.BoolVar(&flagEnableQueryLogging, "enable-query-logging", true, "enable logging of database queries")

	pflag.Parse()

	// Initialize logging.
	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()
	level, err := zerolog.ParseLevel(flagLogLevel)
	if err != nil {
		log.Error().Err(err).Msg("could not parse log level")
		return failure
	}
	log = log.Level(level)

	if flagDatabase == "" {
		log.Error().Msg("database address is required")
		return failure
	}

	// Enable GORM logging if database query logs are enabled.
	var dblog logger.Interface
	if flagEnableQueryLogging {
		dblog = gormzerolog.NewWithLogger(log)
	} else {
		dblog = logger.Default.LogMode(logger.Silent)
	}

	dbCfg := gorm.Config{
		Logger: dblog,
	}
	db, err := gorm.Open(postgres.Open(flagDatabase), &dbCfg)
	if err != nil {
		log.Error().Err(err).Msg("could not connect to database")
		return failure
	}

	storage := storage.New(db)
	apiServer := api.NewServer(storage, log)
	cfg := generated.Config{
		Resolvers: apiServer,
	}

	schema := generated.NewExecutableSchema(cfg)
	gqlServer := handler.NewDefaultServer(schema)

	// Set query complexity limit — each field in a selection set and
	// each nesting level adds the value of one to the overall query
	// complexity.
	if flagComplexityLimit > 0 {
		gqlServer.Use(extension.FixedComplexityLimit(flagComplexityLimit))
	}

	// Initialize Echo Web Server.
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Inject zerolog logger into echo.
	slog := lecho.From(log)
	server.Logger = lecho.From(log)
	server.Use(lecho.Middleware(lecho.Config{Logger: slog}))

	// Initialize server endpoints.
	server.POST(defaultGraphQLEndpoint, echoHandler(gqlServer))

	// If GraphQL Playground is enabled, initialize the handler.
	if flagEnablePlayground {

		// Create a playground handler.
		playground := playground.Handler("GraphQL playground", defaultGraphQLEndpoint)

		// Set the echo handler for the playground.
		server.GET(flagPlayground, echo.WrapHandler(playground))

		// Log the playground URL.
		playgroundURL := formatPlaygroundURL(flagBind, flagPlayground)
		log.Info().Str("address", playgroundURL).Msg("GraphQL playground URL")
	}

	// This section launches the main executing components in their own
	// goroutine, so they can run concurrently. Afterwards, we wait for an
	// interrupt signal in order to proceed with the next section.
	done := make(chan struct{})
	failed := make(chan struct{})
	go func() {
		log.Info().Msg("analytics server starting")
		err := server.Start(flagBind)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Warn().Err(err).Msg("analytics server failed")
			close(failed)
		} else {
			close(done)
		}
		log.Info().Msg("analytics server stopped")
	}()

	select {
	case <-sig:
		log.Info().Msg("analytics server stopping")
	case <-done:
		log.Info().Msg("analytics server done")
	case <-failed:
		log.Warn().Msg("analytics server aborted")
	}

	go func() {
		<-sig
		log.Warn().Msg("forcing exit")
		os.Exit(1)
	}()

	// The following code starts a shut down with a certain timeout and makes
	// sure that the main executing components are shutting down within the
	// allocated shutdown time. Otherwise, we will force the shutdown and log
	// an error. We then wait for shutdown on each component to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("could not shut down analytics server")
		return failure
	}

	return success
}

// Create a formatted link for the GraphQL Playground URL.
func formatPlaygroundURL(address string, path string) string {

	path = strings.TrimPrefix(path, "/")

	if strings.HasPrefix(address, ":") {
		return fmt.Sprintf("http://localhost%s/%s", address, path)
	}

	return fmt.Sprintf("http://%s/%s", address, path)
}

// Create an echo.HandlerFunc for the GraphQL server.
func echoHandler(h *handler.Server) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		h.ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}
