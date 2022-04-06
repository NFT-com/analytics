package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/graph-api/graph/api"
	"github.com/NFT-com/graph-api/graph/models/api"
)

// Collection retrieves a single collection from its ID.
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

// CollectionByContract retrieves a single collection based on the chain ID and the contract address.
func (s *Storage) CollectionByContract(chainID string, contract string) (*api.Collection, error) {

	if chainID == "" {
		return nil, errors.New("chain ID is required")
	}
	if contract == "" {
		return nil, errors.New("contract address is required")
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

// Collections retrieves the list of collections on a chain.
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

// CollectionNFTs retrieves the list of NFTs in a specific collection.
func (s *Storage) CollectionNFTs(collectionID string) ([]*api.NFT, error) {

	var nfts []*api.NFT
	err := s.db.Where(api.NFT{
		Collection: collectionID,
	}).
		Find(&nfts).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil

}

// CollectionsByChain retrieves a list of collections on a specified Chain.
func (s *Storage) CollectionsByChain(chainID string) ([]*api.Collection, error) {

	var collections []*api.Collection
	err := s.db.Where(api.Collection{
		ChainID: chainID,
	}).
		Find(&collections).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
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
