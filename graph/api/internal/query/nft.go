package query

import (
	"context"

	"github.com/NFT-com/analytics/graph/query"
)

// NFTFields describes the paths to the relevant fields in an NFT query.
type NFTFields struct {
	Traits       string
	TraitRarity  string
	Rarity       string
	Owners       string
	Price        string
	AveragePrice string
}

// NFT is used to describe the NFT detail query, namely whether the GraphQL query
// requires fetching of expensive fields such as traits, rarity information or
// aggregated data (stats) for the NFT.
type NFT struct {
	Traits       bool
	TraitRarity  bool
	Rarity       bool
	Owners       bool
	Price        bool
	AveragePrice bool
}

// ParseNFTQuery parses the GraphQL query according to the specified configuration.
func ParseNFTQuery(fields NFTFields, ctx context.Context) *NFT {

	selection := query.GetSelection(ctx)

	query := NFT{
		Traits:       selection.Has(fields.Traits),
		TraitRarity:  selection.Has(fields.TraitRarity),
		Owners:       selection.Has(fields.Owners),
		Rarity:       selection.Has(fields.Rarity),
		Price:        selection.Has(fields.Price),
		AveragePrice: selection.Has(fields.AveragePrice),
	}

	return &query
}

// NeedRarity returns true if either overall rarity or trait rarity is needed.
func (q NFT) NeedRarity() bool {
	return q.TraitRarity || q.Rarity
}
