package api

import (
	"github.com/rs/zerolog"

	"github.com/NFT-com/analytics/aggregate/api/internal/address"
	"github.com/NFT-com/analytics/aggregate/api/internal/currency"
)

// API provides the Aggregation API functionality.
type API struct {
	stats  Stats
	lookup Lookup
	log    zerolog.Logger

	collections  *address.Cache
	marketplaces *address.Cache
	currencies   *currency.Cache
}

// New creates a new API handler.
func New(stats Stats, lookup Lookup, log zerolog.Logger) *API {

	api := API{
		stats:  stats,
		lookup: lookup,
		log:    log,

		collections:  address.NewCache(),
		marketplaces: address.NewCache(),
		currencies:   currency.NewCache(),
	}

	return &api
}
