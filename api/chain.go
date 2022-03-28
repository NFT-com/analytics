package api

import (
	"github.com/NFT-com/graph-api/models/api"
)

// getChain retrieves a chain based on its ID.
func (s *Server) getChain(id string) (*api.Chain, error) {

	chain, err := s.storage.Chain(id)
	if err != nil {
		s.logError(err).Str("id", id).Msg("could not retrieve chain")
		return nil, errRetrieveChainFailed
	}

	return chain, nil
}

// chains returns a list of all chains.
func (s *Server) chains() ([]*api.Chain, error) {

	chains, err := s.storage.Chains()
	if err != nil {
		s.logError(err).Msg("could not retrieve chains")
		return nil, errRetrieveChainFailed
	}

	return chains, nil
}
