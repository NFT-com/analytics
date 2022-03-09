package storage

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// Retrieve a Chain based on the ID.
func (s *Storage) Chain(id string) (*api.Chain, error) {

	chain := api.Chain{
		ID: id,
	}

	err := s.db.First(&chain).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chain: %w", err)
	}

	return &chain, nil
}
