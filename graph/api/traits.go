package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

func (s *Server) nftTraits(nft *api.NFT) ([]*api.Trait, error) {

	// Get trait information for the current NFT.
	traits, err := s.storage.NFTTraits(nft.ID)
	if err != nil {
		s.logError(err).
			Str("nft", nft.ID).
			Msg("could not retrieve NFT traits")
		return nil, errRetrieveTraitsFailed
	}

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
