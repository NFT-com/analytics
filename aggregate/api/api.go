package api

import (
	"fmt"

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

func (a *API) lookupCollection(id string) (identifier.Address, error) {

	address, ok := a.collectionCache[id]
	if ok {
		return address, nil
	}

	address, err := a.lookup.Collection(id)
	if err != nil {
		return identifier.Address{}, fmt.Errorf("could not lookup collection: %w", err)
	}

	// FIXME: Add a mutex to sync this.
	a.collectionCache[id] = address

	return address, nil
}

func (a *API) lookupNFT(id string) (identifier.NFT, error) {

	nft, err := a.lookup.NFT(id)
	if err != nil {
		return identifier.NFT{}, fmt.Errorf("could not lookup NFT: %w", err)
	}

	// FIXME: Add caching. Though, due to the amount of NFTs, it may not be feasible
	// to cache ALL NFT IDs. Instad, perhaps handle them by group.

	return nft, nil
}

func (a *API) lookupMarketplace(id string) ([]identifier.Address, error) {

	addresses, err := a.lookup.Marketplace(id)
	if err != nil {
		return nil, fmt.Errorf("could not lookup marketplace: %w", err)
	}

	// FIXME: Add caching.

	return addresses, nil
}
