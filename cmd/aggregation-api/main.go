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
	"github.com/ziflex/lecho/v2"

	gormzerolog "github.com/wei840222/gorm-zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/NFT-com/analytics/aggregate/api"
	"github.com/NFT-com/analytics/aggregate/lookup"
	"github.com/NFT-com/analytics/aggregate/stats"
)

const (
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
		flagGraphDatabase      string
		flagEventsDatabase     string
		flagLogLevel           string
		flagEnableQueryLogging bool

		flagGraphDBConnections      int
		flagGraphDBIdleConnections  int
		flagEventsDBConnections     int
		flagEventsDBIdleConnections int
	)

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagEventsDatabase, "events-database", "e", "", "events database address")
	pflag.StringVarP(&flagGraphDatabase, "graph-database", "g", "", "graph database address")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")
	pflag.BoolVar(&flagEnableQueryLogging, "enable-query-logging", true, "enable logging of database queries")

	pflag.IntVar(&flagGraphDBConnections, "graph-db-connection-limit", defaultDBMaxConnections, "maximum number of connections to graph database, -1 for unlimited")
	pflag.IntVar(&flagGraphDBIdleConnections, "graph-db-idle-connection-limit", defaultDBIdleConnections, "maximum number of idle connections to graph database")
	pflag.IntVar(&flagEventsDBConnections, "events-db-connection-limit", defaultDBMaxConnections, "maximum number of connections to events database, -1 for unlimited")
	pflag.IntVar(&flagEventsDBIdleConnections, "events-db-idle-connection-limit", defaultDBIdleConnections, "maximum number of idle connections to events database")

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

	// Connect to the Events database.
	dbCfg := gorm.Config{
		Logger: dblog,
	}
	eventsDB, err := gorm.Open(postgres.Open(flagEventsDatabase), &dbCfg)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	// Limit the number of database connections to the events database.
	edb, err := eventsDB.DB()
	if err != nil {
		return fmt.Errorf("could not get database connection for events DB: %w", err)
	}
	edb.SetMaxOpenConns(flagEventsDBConnections)
	edb.SetMaxIdleConns(flagEventsDBIdleConnections)

	// Create stats handler.
	stats := stats.New(eventsDB)

	// Connect to the Graph database.
	graphDB, err := gorm.Open(postgres.Open(flagGraphDatabase), &dbCfg)
	if err != nil {
		return fmt.Errorf("could not connect to graph database: %w", err)
	}
	// Limit the number of database connections to the graph database.
	gdb, err := graphDB.DB()
	if err != nil {
		return fmt.Errorf("could not get database connection for graph DB: %w", err)
	}
	gdb.SetMaxOpenConns(flagGraphDBConnections)
	gdb.SetMaxIdleConns(flagGraphDBIdleConnections)

	// Create lookup handler.
	lookup := lookup.New(graphDB)

	// Create the API.
	api := api.New(stats, lookup, log)

	// Initialize Echo Web Server.
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Inject zerolog logger into echo.
	slog := lecho.From(log)
	server.Logger = lecho.From(log)
	server.Use(lecho.Middleware(lecho.Config{Logger: slog}))

	// Initialize routes.

	// Collection stats - current.
	server.GET("/collections/:id/volume", api.CollectionVolume)
	server.GET("/collections/:id/market_cap", api.CollectionMarketCap)
	server.GET("/collections/:id/sales", api.CollectionSales)
	server.POST("/collections/batch/volume", api.CollectionBatchVolume)
	server.POST("/collections/batch/market_cap", api.CollectionBatchMarketCap)

	// Collection stats - historic.
	server.GET("/collections/:id/volume/history", api.CollectionVolumeHistory)
	server.GET("/collections/:id/market_cap/history", api.CollectionMarketCapHistory)
	server.GET("/collections/:id/sales/history", api.CollectionSalesHistory)
	server.GET("/collections/:id/size/history", api.CollectionSizeHistory)
	server.GET("/collections/:id/average/history", api.CollectionAverageHistory)
	server.GET("/collections/:id/floor/history", api.CollectionFloorHistory)

	// Marketplace stats - current.
	server.GET("/marketplaces/:id/volume", api.MarketplaceVolume)
	server.GET("/marketplaces/:id/market_cap", api.MarketplaceMarketCap)
	server.GET("/marketplaces/:id/sales", api.MarketplaceSales)
	server.GET("/marketplaces/:id/users", api.MarketplaceUsers)

	// Marketplace stats - historic.
	server.GET("/marketplaces/:id/volume/history", api.MarketplaceVolumeHistory)
	server.GET("/marketplaces/:id/market_cap/history", api.MarketplaceMarketCapHistory)
	server.GET("/marketplaces/:id/sales/history", api.MarketplaceSalesHistory)
	server.GET("/marketplaces/:id/users/history", api.MarketplaceUsersHistory)

	// NFT stats - current.
	server.GET("/nfts/:id/price", api.NFTPrice)
	server.POST("/nfts/batch/price", api.NFTBatchPrice)

	// NFT stats - historic.
	server.GET("/nfts/:id/average", api.NFTAveragePrice)
	server.GET("/nfts/:id/price/history", api.NFTPriceHistory)

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
