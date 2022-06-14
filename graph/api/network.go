package api

import (
	"github.com/NFT-com/analytics/graph/models/api"
)

// getNetwork retrieves a network based on its ID.
func (s *Server) getNetwork(id string) (*api.Network, error) {

	network, err := s.storage.Network(id)
	if err != nil {
		s.logError(err).Str("id", id).Msg("could not retrieve network")
		return nil, errRetrieveNetworkFailed
	}

	return network, nil
}

// networks returns all known networks.
func (s *Server) networks() ([]*api.Network, error) {

	networks, err := s.storage.Networks()
	if err != nil {
		s.logError(err).Msg("could not retrieve networks")
		return nil, errRetrieveNetworkFailed
	}

	return networks, nil
}
