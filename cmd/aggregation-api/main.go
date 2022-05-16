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

	"github.com/NFT-com/graph-api/aggregate/api"
	"github.com/NFT-com/graph-api/aggregate/stats"
	"github.com/NFT-com/graph-api/aggregate/storage"
)

// Endpoint paths.
// FIXME: Revisit when implemented
const (
	marketCapEndpoint = "/market_cap"
	volumeEndpoint    = "/volume"
	floorEndpoint     = "/floor"
	averageEndpoint   = "/average"
	salesEndpoint     = "/sale_count"
	nftEndpoint       = "/nft_count"
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
		flagDatabase           string
		flagLogLevel           string
		flagEnableQueryLogging bool
	)

	// FIXME: Add database connection limit.

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagDatabase, "database", "d", "", "database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
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

	stats := stats.New(db)

	storage := storage.New(db)

	api := api.New(stats, storage, log)

	// Initialize Echo Web Server.
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Inject zerolog logger into echo.
	slog := lecho.From(log)
	server.Logger = lecho.From(log)
	server.Use(lecho.Middleware(lecho.Config{Logger: slog}))

	// FIXME: Marketplace will have chainID+address combinations,
	// not a single thing.

	// Initialize routes.
	server.GET("/collection/:id/volume", api.CollectionVolume)
	server.GET("/marketplace/:id/volume", api.MarketplaceVolume)

	server.GET("/collection/:id/market_cap", api.CollectionMarketCap)
	server.GET("/marketplace/:id/market_cap", api.MarketplaceMarketCap)

	server.GET("/collection/:id/sales", api.CollectionSales)
	server.GET("/marketplace/:id/sales", api.MarketplaceSales)

	server.GET("/collection/:id/count", api.CollectionCount)
	server.GET("/marketplace/:id/users", api.MarketplaceUsers)

	server.GET("/collection/:id/average", api.CollectionAverage)
	server.GET("/collection/:id/floor", api.CollectionFloor)

	server.GET("/nft/:id/price", api.NFTPrice)
	server.GET("/nft/:id/average", api.NFTAveragePrice)

	// This section launches the main executing components in their own
	// goroutine, so they can run concurrently. Afterwards, we wait for an
	// interrupt signal in order to proceed with the next section.
	done := make(chan struct{})
	failed := make(chan struct{})
	go func() {
		log.Info().Msg("aggregation API starting")
		err := server.Start(flagBind)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Warn().Err(err).Msg("aggregation API failed")
			close(failed)
		} else {
			close(done)
		}
		log.Info().Msg("aggregation API stopped")
	}()

	select {
	case <-sig:
		log.Info().Msg("aggregation API stopping")
	case <-done:
		log.Info().Msg("aggregation API done")
	case <-failed:
		log.Warn().Msg("aggregation API aborted")
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
		return fmt.Errorf("could not shut down aggregation API server: %w", err)
	}

	return nil
}
