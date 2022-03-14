package api

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// getCollection returns a single collection based on its ID.
func (s *Server) getCollection(id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

// getCollectionByAddress returns a single collection for the specified chain, given its contract address.
func (s *Server) getCollectionByAddress(chainID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByAddress(chainID, contract)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return collection, nil
}

// getCollectionNFTs returns a list of NFTs in a collection.
func (s *Server) getCollectionNFTs(collectionID string) ([]*api.NFT, error) {

	nfts, err := s.storage.CollectionNFTs(collectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return nfts, nil
}

// collections returns a list of collections according to the specified search criteria and sorting options.
func (s *Server) collections(chain *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	collections, err := s.storage.Collections(chain, orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// collectionsByChain returns a list of collections on a given chain.
func (s *Server) collectionsByChain(chainID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByChain(chainID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// CollectionListings returns a list of marketplaces where the collection is listed on.
func (s *Server) collectionsListings(collectionID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesForCollection(collectionID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces for a collection: %w", err)
	}

	return marketplaces, nil
}
