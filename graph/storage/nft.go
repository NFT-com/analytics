package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
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
		Where("LOWER(c.contract_address) = LOWER(?)", contract).
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
func (s *Storage) NFTs(owner *string, collectionID *string, orderBy api.NFTOrder, limit uint, prefetchOwners bool) ([]*api.NFT, error) {

	// By default, we'll query a single table since it's faster.
	db := s.db.
		Table("nfts n").
		Select("*").
		Where("n.deleted != TRUE")

	// If filtering based on the owner is specified, switch to querying using two tables.
	if owner != nil {
		filter := s.db.
			Table("owners o").
			Select("DISTINCT o.nft_id").
			Where("o.owner != ?", identifier.ZeroAddress).
			Where("LOWER(o.owner) = LOWER(?)", *owner).
			Group("LOWER(o.owner), o.nft_id").
			Having("SUM(number) > ?", 0)

		db = s.db.
			Table("nfts n").
			Select("n.*").
			Where("n.id IN (?)", filter).
			Where("n.deleted != TRUE")
	}

	// Set collection filter if specified.
	if collectionID != nil {
		db = db.Where("n.collection_id = ?", collectionID)
	}

	// Set `orderBy` if applicable - only for creation time we can directly use the DB sorting.
	if orderBy.Field == api.NFTOrderFieldCreationTime {
		orderClause := fmt.Sprintf("%s %s", creationTimeColumnName, orderBy.Direction)
		db = db.Order(orderClause)
	}

	if limit > 0 {
		db = db.Limit(int(limit))
	}

	// Execute query.
	var nfts []*api.NFT
	err := db.Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	if !prefetchOwners {
		return nfts, nil
	}

	// Lookup owner for the list of NFTs.
	var ids []string
	for _, nft := range nfts {
		ids = append(ids, nft.ID)
	}

	owners, err := s.nftListOwners(ids)
	if err != nil {
		return nil, fmt.Errorf("could not lookup owners for the NFT list: %w", err)
	}

	for _, nft := range nfts {
		nft.Owners = owners[nft.ID]
	}

	return nfts, nil
}
