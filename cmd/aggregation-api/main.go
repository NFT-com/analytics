package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
)

const (
	// Status codes.
	success = 0
	failure = 1
)

func main() {
	os.Exit(run())
}

func run() int {

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
		log.Error().Err(err).Msg("could not parse log level")
		return failure
	}
	log = log.Level(level)

	log.Error().Msg("TBD: not implemented")
	return failure
}
