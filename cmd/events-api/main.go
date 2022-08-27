package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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

	"github.com/NFT-com/analytics/events/api"
	"github.com/NFT-com/analytics/events/storage"
)

const (
	// Server endpoints.
	transferEndpoint = "/transfers/"
	saleEndpoint     = "/sales/"

	// Default event batch size.
	defaultBatchSize = 100

	defaultDBMaxConnections  = 70
	defaultDBIdleConnections = 20
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
		flagBind               string
		flagBatchSize          uint
		flagDatabase           string
		flagLogLevel           string
		flagEnableQueryLogging bool
		flagDBConnections      int
		flagDBIdleConnections  int
	)

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.UintVarP(&flagBatchSize, "batch-size", "s", defaultBatchSize, "default limit for number of events returned in a single call")
	pflag.StringVarP(&flagDatabase, "database", "d", "", "database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
	pflag.IntVar(&flagDBConnections, "db-connection-limit", defaultDBMaxConnections, "maximum number of database connections, -1 for unlimited")
	pflag.IntVar(&flagDBIdleConnections, "db-idle-connection-limit", defaultDBIdleConnections, "maximum number of idle connections")
	pflag.BoolVar(&flagEnableQueryLogging, "enable-query-logging", true, "enable logging of database queries")

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

	// Initialize storage component.
	storage := storage.New(db, storage.WithBatchSize(flagBatchSize))

	// Initialize the API handler.
	api := api.New(storage, log)

	// Initialize Echo Web Server.
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Inject zerolog logger into echo.
	slog := lecho.From(log)
	server.Logger = lecho.From(log)
	server.Use(lecho.Middleware(lecho.Config{Logger: slog}))

	// Initialize routes.
	server.GET("/health", health)
	server.GET(transferEndpoint, api.Transfer)
	server.GET(saleEndpoint, api.Sale)

	// This section launches the main executing components in their own
	// goroutine, so they can run concurrently. Afterwards, we wait for an
	// interrupt signal in order to proceed with the next section.
	done := make(chan struct{})
	failed := make(chan struct{})
	go func() {
		log.Info().Msg("events API server starting")
		err := server.Start(flagBind)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Warn().Err(err).Msg("events API server failed")
			close(failed)
		} else {
			close(done)
		}
		log.Info().Msg("events API server stopped")
	}()

	select {
	case <-sig:
		log.Info().Msg("events API server stopping")
	case <-done:
		log.Info().Msg("events API server done")
	case <-failed:
		log.Warn().Msg("events API server aborted")
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
		return fmt.Errorf("could not shut down events API server: %w", err)
	}

	return nil
}

// health is an HTTP handler that returns an empty '200 OK' response.
func health(_ echo.Context) error {
	return nil
}
