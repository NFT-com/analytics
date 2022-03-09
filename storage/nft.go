package storage

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// NFT returns a single NFT based on the ID.
func (s *Storage) NFT(id string) (*api.NFT, error) {

	nft := api.NFT{
		ID: id,
	}

	err := s.db.First(&nft).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return &nft, nil
}

func (s *Storage) NFTs() ([]*api.NFT, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
