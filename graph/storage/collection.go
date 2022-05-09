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

// CollectionByContract retrieves a single collection based on the network ID and the contract address.
func (s *Storage) CollectionByContract(networkID string, contract string) (*api.Collection, error) {

	if networkID == "" {
		return nil, errors.New("network ID is required")
	}
	if contract == "" {
		return nil, errors.New("contract address is required")
	}

	var collection api.Collection
	err := s.db.
		Where(api.Collection{
			NetworkID: networkID,
			Address:   contract,
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

// Collections retrieves the list of collections on a network.
func (s *Storage) Collections(networkID *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	query := api.Collection{}
	if networkID != nil {
		query.NetworkID = *networkID
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

	query := api.NFT{
		Collection: collectionID,
	}

	var nfts []*api.NFT
	err := s.db.Where(&query).Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil

}

// CollectionsByNetwork retrieves a list of collections on a specified network.
func (s *Storage) CollectionsByNetwork(networkID string) ([]*api.Collection, error) {

	query := api.Collection{
		NetworkID: networkID,
	}

	var collections []*api.Collection
	err := s.db.Where(&query).Find(&collections).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// CollectionSize returns the number of NFTs in a collection.
func (s *Storage) CollectionSize(collectionID string) (uint, error) {

	query := api.NFT{
		Collection: collectionID,
	}

	var count int64
	err := s.db.Where(&query).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("could not retrieve collection size: %w", err)
	}

	if count < 0 {
		return 0, fmt.Errorf("unexpected collection size (got: %d)", count)
	}

	return uint(count), nil
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
