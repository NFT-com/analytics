package api

import (
	"context"

	"github.com/NFT-com/analytics/graph/models/api"
)

// marketplaceCollections returns a list of collections on a specified marketplace.
func (s *Server) marketplaceCollections(ctx context.Context, marketplaceID string) ([]*api.Collection, error) {

	collections, err := s.storage.MarketplaceCollections(marketplaceID)
	if err != nil {
		s.logError(err).
			Str("marketplace", marketplaceID).
			Msg("could not retrieve collections for marketplace")
		return nil, errRetrieveCollectionFailed
	}

	for _, collection := range collections {
		collection, err = s.expandCollectionDetails(ctx, collection)
		if err != nil {
			s.logError(err).Str("id", collection.ID).Msg("retrieving collection details failed")
			return nil, errRetrieveCollectionFailed
		}
	}

	return collections, nil
}

// marketplacesByNetwork returns a list of marketplaces on a specified network.
func (s *Server) marketplacesByNetwork(networkID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesByNetwork(networkID)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Msg("could not retrieve marketplaces for network")
		return nil, errRetrieveMarketplaceFailed
	}

	return marketplaces, nil
}

// marketplaceNetworks returns a list of networks that collections listed on a marketplace reside on.
func (s *Server) marketplaceNetworks(marketplaceID string) ([]*api.Network, error) {

	networks, err := s.storage.MarketplaceNetworks(marketplaceID)
	if err != nil {
		s.logError(err).
			Str("marketplace", marketplaceID).
			Msg("could not retrieve networks for marketplace")
		return nil, errRetrieveNetworkFailed
	}

	return networks, nil
}
