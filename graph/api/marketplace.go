package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

// marketplaceCollections returns a list of collections on a specified marketplace.
func (s *Server) marketplaceCollections(marketplaceID string) ([]*api.Collection, error) {

	collections, err := s.storage.MarketplaceCollections(marketplaceID)
	if err != nil {
		s.logError(err).
			Str("marketplace", marketplaceID).
			Msg("could not retrieve collections for marketplace")
		return nil, errRetrieveCollectionFailed
	}

	return collections, nil
}

// marketplacesByChain returns a list of marketplaces on a specified chain.
func (s *Server) marketplacesByChain(chainID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesByChain(chainID)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Msg("could not retrieve marketplaces for chain")
		return nil, errRetrieveMarketplaceFailed
	}

	return marketplaces, nil
}

// marketplaceChains returns a list of chains that collections listed on a marketplace reside on.
func (s *Server) marketplaceChains(marketplaceID string) ([]*api.Chain, error) {

	chains, err := s.storage.MarketplaceChains(marketplaceID)
	if err != nil {
		s.logError(err).
			Str("marketplace", marketplaceID).
			Msg("could not retrieve chains for marketplace")
		return nil, errRetrieveChainFailed
	}

	return chains, nil
}
