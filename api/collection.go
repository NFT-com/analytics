package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// GetCollection returns a single collection based on its ID.
func (s *Server) GetCollection(id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

// GetCollectionByAddress returns a single collection for the specified chain, given its contract address.
func (s *Server) GetCollectionByAddress(chainID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByAddress(chainID, contract)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

// GetCollectionNFTs returns a list of NFTs in a collection.
func (s *Server) GetCollectionNFTs(collectionID string) ([]*api.NFT, error) {

	nfts, err := s.storage.CollectionNFTs(collectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return nfts, nil
}

// Collections returns a list of collections according to the specified search criteria and sorting options.
func (s *Server) Collections(chain *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	collections, err := s.storage.Collections(chain, orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// CollectionsByChain returns a list of collections on a given chain.
func (s *Server) CollectionsByChain(chainID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByChain(chainID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// CollectionListings returns a list of marketplaces where the collection is listed on.
func (s *Server) CollectionsListings(collectionID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesForCollection(collectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces for a collection: %w", err)
	}

	return marketplaces, nil
}
