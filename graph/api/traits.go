package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

func (s *Server) nftTraits(nftID string) ([]*api.TraitRatio, error) {

	traits, err := s.storage.NFTTraitRatio(nftID)
	if err != nil {
		s.logError(err).
			Str("nft", nftID).
			Msg("could not retrieve NFT traits")
		return nil, errRetrieveTraitsFailed
	}

	return traits, nil
}
