package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/indexer-api/api"
	"github.com/NFT-com/indexer-api/models/api"
)

// Collection will retrieve a single collection based on the ID.
func (s *Storage) Collection(id string) (*api.Collection, error) {

	collection := api.Collection{
		ID: id,
	}

	err := s.db.First(&collection).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server.ErrRecordNotFound
	}
	if err != nil {
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server.ErrRecordNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return &collection, nil
}

func (s *Storage) Collections(chain *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	query := api.Collection{}
	if chain != nil {
		query.ChainID = *chain
	}
	db := s.db.Where(query)

	orderClause, err := formatCollectionOrderBy(orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not prepare order clause: %w", err)
	}

	db = db.Order(orderClause)

	var collections []*api.Collection
	err = db.Find(&collections).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
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

func formatCollectionOrderBy(clause api.CollectionOrder) (string, error) {

	var field string

	switch clause.Field {

	case api.CollectionOrderFieldCreationTime:
		field = creationTimeColumnName

	// FIXME: Remove when sorting by value becomes possible.
	default:
		return "", errors.New("unsupported sorting option")
	}

	formatted := fmt.Sprintf("%s %s", field, clause.Direction)

	return formatted, nil
}
