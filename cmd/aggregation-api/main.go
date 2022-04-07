package main

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func run() error {

	var (
		flagBind      string
		flagEventsAPI string
		flagLogLevel  string
	)

	pflag.StringVarP(&flagBind, "bind", "b", ":8080", "bind address for serving requests")
	pflag.StringVarP(&flagEventsAPI, "events-api", "e", "", "URL of the Events API")
	pflag.StringVarP(&flagLogLevel, "log-level", "l", "info", "log level")

	pflag.Parse()

	// Initialize logging.
	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	level, err := zerolog.ParseLevel(flagLogLevel)
	if err != nil {
		return errors.New("could not parse log level")
	}
	log = log.Level(level)
	zerolog.SetGlobalLevel(level)

	return errors.New("TBD: not implemented")
}
