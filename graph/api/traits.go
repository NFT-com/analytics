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
