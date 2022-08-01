package query

import (
	"context"

	"github.com/NFT-com/analytics/graph/query"
)

// CollectionFields describes the paths to the relevant fields in a collection query.
type CollectionFields struct {
	Volume       string
	MarketCap    string
	Sales        string
	NFTs         string
	StartCursor  string
	NFT          NFTFields
	NFTArguments CollectionNFTArguments
}

// CollectionNFTArguments describes the paging arguments for the collection NFT list.
type CollectionNFTArguments struct {
	First string
	After string
}

// Collection describes the collection query, specifically whether the GraphQL query requires
// retrieving NFTs or fetching aggregated fields (stats) for the collection.
type Collection struct {
	Volume      bool
	MarketCap   bool
	Sales       bool
	NFTs        bool
	StartCursor bool
	NFT         CollectionNFTs
}

// CollectionNFTs describes the NFT query context from within the Collection object.
type CollectionNFTs struct {
	Fields      NFT
	StartCursor bool
	Arguments   NFTArguments
}

// NFTArguments represents the arguments for the NFT field of a Collection GraphQL query.
type NFTArguments struct {
	First uint
	After string
}

// ParseCollectionQuery parses the GraphQL query according to the specified configuration.
func ParseCollectionQuery(fields CollectionFields, ctx context.Context) *Collection {

	selection := query.GetSelection(ctx)

	// Paging arguments.
	args := selection.Args(fields.NFTs)

	// NOTE: The GraphQL framework will already handle the field type validation.
	first, _ := args[fields.NFTArguments.First].(int64)
	after, _ := args[fields.NFTArguments.After].(string)

	// NFT selection.
	nft := CollectionNFTs{
		Fields: NFT{
			Rarity:       selection.Has(fields.NFT.Rarity),
			Traits:       selection.Has(fields.NFT.Traits),
			TraitRarity:  selection.Has(fields.NFT.TraitRarity),
			Owners:       selection.Has(fields.NFT.Owners),
			Price:        selection.Has(fields.NFT.Price),
			AveragePrice: selection.Has(fields.NFT.AveragePrice),
		},
		Arguments: NFTArguments{
			First: uint(first),
			After: after,
		},
		StartCursor: selection.Has(fields.StartCursor),
	}

	query := Collection{
		Volume:      selection.Has(fields.Volume),
		MarketCap:   selection.Has(fields.MarketCap),
		Sales:       selection.Has(fields.Sales),
		NFTs:        selection.Has(fields.NFTs),
		StartCursor: selection.Has(fields.StartCursor),
		NFT:         nft,
	}
	return &query
}

// NeedStats returns true if any of the aggregated data is required.
func (c *Collection) NeedStats() bool {
	return c.Volume || c.MarketCap || c.Sales
}
