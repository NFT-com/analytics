package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

const (
	traitRarityField = "rarity"
)

func (s *Server) nftTraits(nft *api.NFT, wantRarity bool) ([]*api.Trait, error) {

	// Get trait information for the current NFT.
	traits, err := s.storage.NFTTraits(nft.ID, wantRarity)
	if err != nil {
		s.logError(err).
			Str("nft", nft.ID).
			Msg("could not retrieve NFT traits")
		return nil, errRetrieveTraitsFailed
	}

	// If we don't need to know the NFT rarity, we're done.
	if !wantRarity {
		return traits, nil
	}

	// Find rarity for missing traits.

	foundTraits := make([]string, 0, len(traits))
	for _, trait := range traits {
		foundTraits = append(foundTraits, trait.Type)
	}

	missingTraits, err := s.storage.NFTMissingTraitRatio(nft.Collection, foundTraits)
	if err != nil {
		s.logError(err).
			Str("nft", nft.ID).
			Msg("could not retrieve NFT traits")
		return nil, errRetrieveTraitsFailed
	}

	traits = append(traits, missingTraits...)

	return traits, nil
}

func calcRarity(traits []*api.Trait) float64 {

	// Calculate rarity of an NFT by multiplying the ratios of individual traits.
	// For example, if NFT has two traits that are present in 50% of
	// NFTs in a collection, the rarity is 0.5*0.5 = 0.25.
	rarity := 1.0
	for _, trait := range traits {
		rarity = rarity * trait.Rarity
	}

	return rarity
}
