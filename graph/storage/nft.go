package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/analytics/graph/api"
	"github.com/NFT-com/analytics/graph/models/api"
)

const (
	creationTimeColumnName = "created_at"
)

// NFT retrieves a single NFT based on the ID.
func (s *Storage) NFT(id string) (*api.NFT, error) {

	nft := api.NFT{
		ID: id,
	}

	err := s.db.First(&nft).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server.ErrRecordNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return &nft, nil
}

// NFTByTokenID retrieves a single NFT based on the network, contract and the tokenID.
func (s *Storage) NFTByTokenID(networkID string, contract string, tokenID string) (*api.NFT, error) {

	if networkID == "" {
		return nil, errors.New("network ID is required")
	}
	if contract == "" {
		return nil, errors.New("contract address is required")
	}
	if tokenID == "" {
		return nil, errors.New("token ID is required")
	}

	var nft api.NFT
	err := s.db.
		Joins("INNER JOIN collections c ON collection_id = c.id").
		Where("c.network_id = ?", networkID).
		Where("c.contract_address = ?", contract).
		Where("token_id = ?", tokenID).
		First(&nft).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, server.ErrRecordNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return &nft, nil
}

// NFTs retrieves a list of NFTs fitting the specified criteria.
func (s *Storage) NFTs(owner *string, collectionID *string, orderBy api.NFTOrder, limit uint) ([]*api.NFT, error) {

	// Apply explicit query filters - the token owner and collection ID.
	query := api.NFT{}

	// FIXME: Reintroduce owner handling.

	if collectionID != nil {
		query.Collection = *collectionID
	}
	db := s.db.Where(query)

	// Set `orderBy` if applicable - only for creation time we can directly use the DB sorting.
	if orderBy.Field == api.NFTOrderFieldCreationTime {
		orderClause := fmt.Sprintf("%s %s", creationTimeColumnName, orderBy.Direction)
		db = db.Order(orderClause)
	}

	if limit > 0 {
		db = db.Limit(int(limit))
	}

	var nfts []*api.NFT
	err := db.Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}
