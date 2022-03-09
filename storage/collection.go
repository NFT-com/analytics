package storage

import (
	"errors"
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// Collection will retrieve a single collection based on the ID.
func (s *Storage) Collection(id string) (*api.Collection, error) {

	collection := api.Collection{
		ID: id,
	}

	err := s.db.First(&collection).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return &collection, nil
}

// CollectionByAddresss returns a single collection based on the chain ID and the contract addresss.
func (s *Storage) CollectionByAddress(chainID string, contract string) (*api.Collection, error) {

	if chainID == "" || contract == "" {
		return nil, errors.New("mandatory fields missing")
	}

	var collection api.Collection
	err := s.db.
		Where(api.Collection{
			ChainID: chainID,
			Address: contract,
		}).
		First(&collection).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return &collection, nil
}

func (s *Storage) Collections() ([]*api.Collection, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}

// CollectionNFTs will return the list of NFTs in a specific Collection.
func (s *Storage) CollectionNFTs(collectionID string) ([]*api.NFT, error) {

	var nfts []*api.NFT
	err := s.db.Where(api.NFT{
		CollectionID: collectionID,
	}).
		Find(&nfts).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil

}

// Retrieve a list of Cllections on a specified Chain.
func (s *Storage) CollectionsByChain(chainID string) ([]*api.Collection, error) {

	var collections []*api.Collection
	err := s.db.Where(api.Collection{
		ChainID: chainID,
	}).
		Find(&collections).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %W", err)
	}

	return collections, nil
}
