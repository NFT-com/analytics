package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// marketplaceCollections returns a list of collections on a specified marketplace.
func (s *Server) marketplaceCollections(marketplaceID string) ([]*api.Collection, error) {

	collections, err := s.storage.MarketplaceCollections(marketplaceID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections on a marketplace: %w", err)
	}

	return collections, nil
}

// marketplacesByChain returns a list of marketplaces on a specified chain.
func (s *Server) marketplacesByChain(chainID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesByChain(chainID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// marketplaceChains returns a list of chains that collections listed on a marketplace reside on.
func (s *Server) marketplaceChains(marketplaceID string) ([]*api.Chain, error) {

	chains, err := s.storage.MarketplaceChains(marketplaceID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chains: %w", err)
	}

	return chains, nil
}
