package api

import (
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// Lookup provides collection and marketplace address lookup based on the ID.
type Lookup interface {
	// Lookup single entity.
	Collection(id string) (identifier.Address, error)
	Marketplace(id string) ([]identifier.Address, error)
	NFT(id string) (identifier.NFT, error)

	// Lookup batches of IDs.
	NFTs(ids []string) (map[string]identifier.NFT, error)
	Collections(id []string) (map[string]identifier.Address, error)
}
