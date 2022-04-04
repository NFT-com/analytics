package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/graph-api/graph-api/api"
	"github.com/NFT-com/graph-api/graph-api/models/api"
)

// Chain retrieves a single chain based on the ID.
func (s *Storage) Chain(id string) (*api.Chain, error) {

	chain := api.Chain{
		ID: id,
	}

	err := s.db.First(&chain).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server.ErrRecordNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chain: %w", err)
	}

	return &chain, nil
}

// Chains retrieves a list of all known chains.
func (s *Storage) Chains() ([]*api.Chain, error) {

	var chains []*api.Chain
	err := s.db.Find(&chains).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chains: %w", err)
	}

	return chains, nil
}
