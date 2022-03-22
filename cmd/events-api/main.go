package main

import (
	"os"
	"time"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/storage"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	gormzerolog "github.com/wei840222/gorm-zerolog"
	"github.com/ziflex/lecho/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	// Status codes.
	success = 0
	failure = 1
)

const (
	// Server endpoints.
	mintEndpoint     = "/mints/"
	transferEndpoint = "/transfers/"
	burnEndpoint     = "/burns/"
	saleEndpoint     = "/sales/"
)

func main() {
	os.Exit(run())
}

func run() int {

	var (
		flagBind               string
		flagDatabase           string
		flagLogLevel           string
		flagEnableQueryLogging bool
	)

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagDatabase, "database", "d", "", "database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
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

	// Initialize storage component.
	storage := storage.New(db)

	// Initialize the API handler.
	api := api.New(storage)

	// FIXME: Remove when storage starts getting used.
	_ = storage

	// Initialize Echo Web Server.
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Inject zerolog logger into echo.
	slog := lecho.From(log)
	server.Logger = lecho.From(log)
	server.Use(lecho.Middleware(lecho.Config{Logger: slog}))

	// Initialize routes.
	server.GET(mintEndpoint, api.Mint)
	server.GET(transferEndpoint, api.Transfer)
	server.GET(saleEndpoint, api.Sale)
	server.GET(burnEndpoint, api.Burn)

	return success
}
