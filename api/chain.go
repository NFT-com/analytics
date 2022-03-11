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
