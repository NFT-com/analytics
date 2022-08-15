package api

import (
	"github.com/rs/zerolog"
)

// API provides the Aggregation API functionality.
type API struct {
	stats  Stats
	lookup Lookup
	log    zerolog.Logger

	collections  *addressCache
	marketplaces *addressCache
	currencies   *currencyCache
}

// New creates a new API handler.
func New(stats Stats, lookup Lookup, log zerolog.Logger) *API {

	api := API{
		stats:  stats,
		lookup: lookup,
		log:    log,

		collections:  newAddressCache(),
		marketplaces: newAddressCache(),
	}

	return &api
}
