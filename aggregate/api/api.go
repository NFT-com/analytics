package api

import (
	"github.com/rs/zerolog"
)

// API provides the Aggregation API functionality.
type API struct {
	stats Stats
	log   zerolog.Logger
}

// New creates a new API handler.
func New(stats Stats, log zerolog.Logger) *API {

	api := API{
		stats: stats,
		log:   log,
	}

	return &api
}
