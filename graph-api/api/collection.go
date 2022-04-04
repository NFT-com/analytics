package api

import (
	"github.com/NFT-com/graph-api/graph-api/models/api"
)

// getCollection returns a single collection based on its ID.
func (s *Server) getCollection(id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	return collection, nil
}

// getCollectionByContract returns a single collection for the specified chain, given its contract address.
func (s *Server) getCollectionByContract(chainID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByContract(chainID, contract)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Str("contract", contract).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	return collection, nil
}

// getCollectionNFTs returns a list of NFTs in a collection.
func (s *Server) getCollectionNFTs(collectionID string) ([]*api.NFT, error) {

	nfts, err := s.storage.CollectionNFTs(collectionID)
	if err != nil {
		s.logError(err).
			Str("id", collectionID).
			Msg("could not retrieve NFTs for a collection")
		return nil, errRetrieveNFTFailed
	}

	return nfts, nil
}

// collections returns a list of collections according to the specified search criteria and sorting options.
func (s *Server) collections(chain *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	collections, err := s.storage.Collections(chain, orderBy)
	if err != nil {
		log := s.logError(err)
		if chain != nil {
			log = log.Str("chain", *chain)
		}
		log.Msg("could not retrieve collections")
		return nil, errRetrieveCollectionFailed
	}

	return collections, nil
}

// collectionsByChain returns a list of collections on a given chain.
func (s *Server) collectionsByChain(chainID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByChain(chainID)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Msg("could not retrieve collections for a chain")
		return nil, errRetrieveCollectionFailed
	}

	return collections, nil
}

// CollectionListings returns a list of marketplaces where the collection is listed on.
func (s *Server) collectionsListings(collectionID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesForCollection(collectionID)
	if err != nil {
		s.logError(err).
			Str("collection", collectionID).
			Msg("could not retrieve marketplaces for a collection")
		return nil, errRetrieveMarketplaceFailed
	}

	return marketplaces, nil
}
