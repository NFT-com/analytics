package api

import (
	"fmt"

	"github.com/rs/zerolog"
)

// API provides the Aggregation API functionality.
type API struct {
	stats  Stats
	lookup Lookup
	log    zerolog.Logger

	collectionCache map[string]collectionAddress
}

// New creates a new API handler.
func New(stats Stats, lookup Lookup, log zerolog.Logger) *API {

	api := API{
		stats:  stats,
		lookup: lookup,
		log:    log,
	}

	return &api
}

func (a *API) lookupCollection(id string) (uint, string, error) {

	address, ok := a.collectionCache[id]
	if ok {
		return address.chainID, address.contractAddress, nil
	}

	chainID, contractAddress, err := a.lookup.CollectionAddress(id)
	if err != nil {
		return 0, "", fmt.Errorf("could not lookup collection: %w", err)
	}

	address = collectionAddress{
		chainID:         chainID,
		contractAddress: contractAddress,
	}

	// FIXME: Add a mutex to sync this.
	a.collectionCache[id] = address

	return chainID, contractAddress, nil
}
