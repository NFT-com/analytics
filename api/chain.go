package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// GetChain retrieves a chain based on its ID.
func (s *Server) GetChain(id string) (*api.Chain, error) {

	chain, err := s.storage.Chain(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chain: %w", err)
	}

	return chain, nil
}

// Chains returns a list of all chains.
func (s *Server) Chains() ([]*api.Chain, error) {

	chains, err := s.storage.Chains()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chains: %w", err)
	}

	return chains, nil
}
