package api

import (
	"context"
	"fmt"

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

	// Parse the collection query.
	query := parseCollectionQuery(ctx)

	for _, collection := range collections {
		// Get collection details, as needed.
		err = s.expandCollectionDetails(query, collection)
		if err != nil {
			return nil, fmt.Errorf("could not expand collection details (id: %v): %w", collection.ID, err)
		}
	}

	return collections, nil
}

// marketplacesByNetwork returns a list of marketplaces on a specified network.
func (s *Server) marketplacesByNetwork(ctx context.Context, networkID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesByNetwork(networkID)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Msg("could not retrieve marketplaces for network")
		return nil, errRetrieveMarketplaceFailed
	}

	// Parse the collection query to see if we need any stats.
	query := parseMarketplaceQuery(ctx)

	// If we don't need any stats, just return the data we have.
	if !query.NeedStats() {
		return marketplaces, nil
	}

	// Retrieve any statistics from the aggregation API.
	for _, marketplace := range marketplaces {
		err = s.expandMarketplaceStats(query, marketplace)
		if err != nil {
			// Continue even if stats could not be retrieved (e.g. API was unavailable).
			s.log.Error().
				Err(err).
				Str("id", marketplace.ID).
				Msg("could not retrieve marketplace stats")
		}
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
