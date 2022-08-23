package api

import (
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

type SingleEntityLookup interface {
	Collection(id string) (identifier.Address, error)
	Marketplace(id string) ([]identifier.Address, error)
	NFT(id string) (identifier.NFT, error)
	CurrencyID(currency identifier.Currency) (string, error)
}

type BatchEntityLookup interface {
	Collections(id []string) (map[string]identifier.Address, error)
	CollectionNFTs(id string) (map[string]identifier.NFT, error)
}

// Lookup provides collection and marketplace address lookup based on the ID.
type Lookup interface {
	SingleEntityLookup
	BatchEntityLookup
}
