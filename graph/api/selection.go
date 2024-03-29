package api

import (
	"context"

	"github.com/NFT-com/analytics/graph/api/internal/query"
	gql "github.com/NFT-com/analytics/graph/query"
)

// parseNFTQuery parses the GraphQL NFT query.
func parseNFTQuery(ctx context.Context) *query.NFT {

	paths := query.NFTFields{
		Traits:       gql.FieldPath(fieldTraits),
		TraitRarity:  gql.FieldPath(fieldTraits, fieldRarity),
		Rarity:       gql.FieldPath(fieldRarity),
		Owners:       gql.FieldPath(fieldOwners),
		Price:        gql.FieldPath(fieldPrice),
		AveragePrice: gql.FieldPath(fieldAveragePrice),
	}

	return query.ParseNFTQuery(paths, ctx)
}

// parseCollectionQuery parses the GraphQL Collection query.
func parseCollectionQuery(ctx context.Context) *query.Collection {

	paths := query.CollectionFields{
		Volume:      gql.FieldPath(fieldVolume),
		MarketCap:   gql.FieldPath(fieldMarketCap),
		Sales:       gql.FieldPath(fieldSales),
		NFTs:        gql.FieldPath(fieldNFTs),
		StartCursor: gql.FieldPath(fieldNFTs, fieldPageInfo, fieldStartCursor),
		NFT: query.NFTFields{
			Traits:       gql.FieldPath(fieldNFTs, fieldEdges, fieldNode, fieldTraits),
			TraitRarity:  gql.FieldPath(fieldNFTs, fieldEdges, fieldNode, fieldTraits, fieldRarity),
			Rarity:       gql.FieldPath(fieldNFTs, fieldEdges, fieldNode, fieldRarity),
			Owners:       gql.FieldPath(fieldNFTs, fieldEdges, fieldNode, fieldOwners),
			Price:        gql.FieldPath(fieldNFTs, fieldEdges, fieldNode, fieldPrice),
			AveragePrice: gql.FieldPath(fieldNFTs, fieldEdges, fieldNode, fieldAveragePrice),
		},
		NFTArguments: query.CollectionNFTArguments{
			First: argumentFirst,
			After: argumentAfter,
		},
	}

	return query.ParseCollectionQuery(paths, ctx)
}

// parseMarketplaceQuery parses the GraphQL Marketplace query.
func parseMarketplaceQuery(ctx context.Context) *query.Marketplace {

	paths := query.MarketplaceFields{
		Volume:    gql.FieldPath(fieldVolume),
		MarketCap: gql.FieldPath(fieldMarketCap),
		Sales:     gql.FieldPath(fieldSales),
		Users:     gql.FieldPath(fieldUsers),
	}

	return query.ParseMarketplaceQuery(paths, ctx)
}
