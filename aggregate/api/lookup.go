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
	// FIXME: Consider having this as the only option.
	NFTs(id []string) (map[string]identifier.NFT, error)
}
