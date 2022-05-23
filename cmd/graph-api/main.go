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
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	gormzerolog "github.com/wei840222/gorm-zerolog"
	"github.com/ziflex/lecho/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/NFT-com/analytics/graph/api"
	"github.com/NFT-com/analytics/graph/generated"
	"github.com/NFT-com/analytics/graph/storage"
)

const (
	defaultPlaygroundPath  = "/"
	defaultGraphQLEndpoint = "/graphql"

	defaultDBMaxConnections  = 70
	defaultDBIdleConnections = 20
	defaultNFTSearchLimit    = 1000
)

func main() {
	err := run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func run() error {

	// Signal catching for clean shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	var (
		flagAggregationAPI     string
		flagBind               string
		flagDatabase           string
		flagLogLevel           string
		flagDBConnections      int
		flagDBIdleConnections  int
		flagEnableQueryLogging bool
		flagEnablePlayground   bool
		flagPlayground         string
		flagComplexityLimit    int
		flagSearchLimit        uint
	)

	pflag.StringVarP(&flagAggregationAPI, "aggregation-api", "a", "", "URL of the Aggregation API")
	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagDatabase, "database", "d", "", "database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
	pflag.IntVar(&flagDBConnections, "db-connection-limit", defaultDBMaxConnections, "maximum number of database connections, -1 for unlimited")
	pflag.IntVar(&flagDBIdleConnections, "db-idle-connection-limit", defaultDBIdleConnections, "maximum number of idle connections")
	pflag.BoolVar(&flagEnablePlayground, "enable-playground", false, "enable GraphQL playground")
	pflag.BoolVar(&flagEnableQueryLogging, "enable-query-logging", true, "enable logging of database queries")
	pflag.StringVar(&flagPlayground, "playground-path", defaultPlaygroundPath, "path for GraphQL playground")
	pflag.IntVar(&flagComplexityLimit, "query-complexity", 0, "GraphQL query complexity limit")
	pflag.UintVar(&flagSearchLimit, "search-limit", defaultNFTSearchLimit, "maximum number of results returned from the NFT search query")

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

	if flagDatabase == "" {
		return errors.New("database address is required")
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
		return fmt.Errorf("could not connect to database: %w", err)
	}
	// Limit the number of database connections.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("could not get database connection: %w", err)
	}
	sqlDB.SetMaxOpenConns(flagDBConnections)
	sqlDB.SetMaxIdleConns(flagDBIdleConnections)

	storage := storage.New(db)
	apiServer := api.NewServer(
		log,
		storage,
		api.WithSearchLimit(flagSearchLimit),
	)
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
		return fmt.Errorf("could not shut down analytics server: %w", err)
	}

	return nil
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
