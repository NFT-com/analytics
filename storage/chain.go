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

// Chains returns a list of known chains.
func (s *Storage) Chains() ([]*api.Chain, error) {

	var chains []*api.Chain
	err := s.db.Find(&chains).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chains: %w", err)
	}

	return chains, nil
}
