package api

import (
	"github.com/rs/zerolog"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// API provides the Aggregation API functionality.
type API struct {
	stats  Stats
	lookup Lookup
	log    zerolog.Logger

	collectionCache map[string]identifier.Address
}

// New creates a new API handler.
func New(stats Stats, lookup Lookup, log zerolog.Logger) *API {

	api := API{
		stats:  stats,
		lookup: lookup,
		log:    log,

		collectionCache: make(map[string]identifier.Address),
	}

	return &api
}
