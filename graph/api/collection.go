package api

import (
	"context"
	"fmt"

	"github.com/NFT-com/analytics/graph/models/api"
)

// getCollection returns a single collection based on its ID.
func (s *Server) getCollection(ctx context.Context, id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	// Parse the collection query.
	query := parseCollectionQuery(ctx)

	// Get collection details, as needed.
	err = s.expandCollectionDetails(query, collection)
	if err != nil {
		return nil, fmt.Errorf("could not expand collection details: %w", err)
	}

	return collection, nil
}

// getCollectionByContract returns a single collection for the specified network, given its contract address.
func (s *Server) getCollectionByContract(ctx context.Context, networkID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByContract(networkID, contract)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Str("contract", contract).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	// Parse the collection query.
	query := parseCollectionQuery(ctx)

	err = s.expandCollectionDetails(query, collection)
	if err != nil {
		return nil, fmt.Errorf("could not expand collection details: %w", err)
	}

	return collection, nil
}

// getCollectionNFTs returns a list of NFTs in a collection.
func (s *Server) getCollectionNFTs(collectionID string, limit uint, afterID string) ([]*api.NFT, bool, error) {

	nfts, lastPage, err := s.storage.CollectionNFTs(collectionID, limit, afterID)
	if err != nil {
		s.logError(err).
			Str("id", collectionID).
			Msg("could not retrieve NFTs for a collection")
		return nil, false, errRetrieveNFTFailed
	}

	return nfts, lastPage, nil
}

// collections returns a list of collections according to the specified search criteria and sorting options.
func (s *Server) collections(ctx context.Context, network *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	collections, err := s.storage.Collections(network, orderBy)
	if err != nil {
		log := s.logError(err)
		if network != nil {
			log = log.Str("network", *network)
		}
		log.Msg("could not retrieve collections")
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

// collectionsByNetwork returns a list of collections on a given network.
func (s *Server) collectionsByNetwork(ctx context.Context, networkID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByNetwork(networkID)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Msg("could not retrieve collections for a network")
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

// collectionsListings returns a list of marketplaces where the collection is listed on.
func (s *Server) collectionsListings(ctx context.Context, collectionID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesForCollection(collectionID)
	if err != nil {
		s.logError(err).
			Str("collection", collectionID).
			Msg("could not retrieve marketplaces for a collection")
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
