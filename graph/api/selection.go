package api

import (
	"context"

	"github.com/NFT-com/analytics/graph/query"
)

type nftQueryConfig struct {
	traitPath       string
	traitRarityPath string
	rarityPath      string
	ownersPath      string
}

// nftQuery is used to describe the NFT detail query, namely whether the GraphQL query
// requires fetching of NFT traits, calculating trait rarity or overall NFT rarity.
type nftQuery struct {
	traits      bool
	traitRarity bool
	nftRarity   bool
	owners      bool
}

// parseNFTQuery parses the GraphQL query with the default configuration.
func parseNFTQuery(ctx context.Context) *nftQuery {
	cfg := nftQueryConfig{
		traitPath:       query.FieldPath(fieldTraits),
		traitRarityPath: query.FieldPath(fieldTraits, fieldRarity),
		rarityPath:      query.FieldPath(fieldRarity),
		ownersPath:      query.FieldPath(fieldOwners),
	}
	return parseNFTQueryWithConfig(cfg, ctx)
}

// parseNFTQueryWithConfig parses the GraphQL query according to the specified configuration.
func parseNFTQueryWithConfig(cfg nftQueryConfig, ctx context.Context) *nftQuery {

	selection := query.GetSelection(ctx)

	query := nftQuery{
		traits:      selection.Has(cfg.traitPath),
		traitRarity: selection.Has(cfg.traitRarityPath),
		nftRarity:   selection.Has(cfg.rarityPath),
		owners:      selection.Has(cfg.ownersPath),
	}

	return &query
}

func (q nftQuery) needRarity() bool {
	return q.traitRarity || q.nftRarity
}
