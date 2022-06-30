package query

import (
	"context"

	"github.com/NFT-com/analytics/graph/query"
)

// CollectionFields describes the paths to the relevant fields in a collection query.
type CollectionFields struct {
	Volume    string
	MarketCap string
	Sales     string
	NFTs      string
	NFT       NFTFields
}

// Collection describes the collection query, specifically whether the GraphQL query requires
// retrieving NFTs or fetching aggregated fields (stats) for the collection.
type Collection struct {
	Volume    bool
	MarketCap bool
	Sales     bool
	NFTs      bool
	NFT       NFT
}

// ParseCollectionQuery parses the GraphQL query according to the specified configuration.
func ParseCollectionQuery(fields CollectionFields, ctx context.Context) *Collection {

	selection := query.GetSelection(ctx)

	query := Collection{
		Volume:    selection.Has(fields.Volume),
		MarketCap: selection.Has(fields.MarketCap),
		Sales:     selection.Has(fields.Sales),
		NFTs:      selection.Has(fields.NFTs),
		NFT: NFT{
			Rarity:       selection.Has(fields.NFT.Rarity),
			Traits:       selection.Has(fields.NFT.Traits),
			TraitRarity:  selection.Has(fields.NFT.TraitRarity),
			Owners:       selection.Has(fields.NFT.Owners),
			Price:        selection.Has(fields.NFT.Price),
			AveragePrice: selection.Has(fields.NFT.AveragePrice),
		},
	}

	return &query
}

// NeedStats returns true if any of the aggregated data is required.
func (c *Collection) NeedStats() bool {
	return c.Volume || c.MarketCap || c.Sales
}
