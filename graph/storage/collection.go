package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/analytics/graph/api"
	"github.com/NFT-com/analytics/graph/models/api"
)

const (
	startHeightColumnName = "start_height"
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
		Where("network_id = ?", networkID).
		Where("LOWER(contract_address) = LOWER(?)", contract).
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

// CollectionNFTs retrieves the list of NFTs in a specific collection, as well as a boolean indicating if there
// are more results or not.
func (s *Storage) CollectionNFTs(collectionID string, limit uint, afterID string) ([]*api.NFT, bool, error) {

	query := s.db.
		Table("nfts").
		Where("collection_id = ?", collectionID).
		Order("id ASC")

	// Request one NFT more than needed, so we know if there are more NFTs after this group.
	if limit > 0 {
		query = query.Limit(int(limit) + 1)
	}

	if afterID != "" {
		query = query.Where("id > ?", afterID)
	}

	var nfts []*api.NFT
	err := query.Find(&nfts).Error
	if err != nil {
		return nil, false, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	// If we requested all tokens, just return all NFTs and notify there are
	// no more left.
	if limit == 0 {
		return nfts, false, nil
	}

	// If the number of returned items is larger than `limit``,
	// there are more NFTs.
	more := uint(len(nfts)) > limit

	// Trim the list to correct size if we have more records.
	if uint(len(nfts)) > limit {
		nfts = nfts[:limit]
	}

	return nfts, more, nil

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
	err := s.db.Model(&api.NFT{}).Where(&query).Count(&count).Error
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
		field = startHeightColumnName

	// FIXME: Remove when sorting by value becomes possible.
	default:
		return "", errors.New("unsupported sorting option")
	}

	formatted := fmt.Sprintf("%s %s", field, clause.Direction)

	return formatted, nil
}
