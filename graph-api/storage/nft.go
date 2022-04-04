package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	server "github.com/NFT-com/graph-api/api"
	"github.com/NFT-com/graph-api/models/api"
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

// NFTByTokenID retrieves a single NFT based on the chain, contract and the tokenID.
func (s *Storage) NFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	if chainID == "" {
		return nil, errors.New("chain ID is required")
	}
	if contract == "" {
		return nil, errors.New("contract address is required")
	}
	if tokenID == "" {
		return nil, errors.New("token ID is required")
	}

	var nft api.NFT
	err := s.db.
		Joins("INNER JOIN collections c ON collection = c.id").
		Where("c.chain_id = ?", chainID).
		Where("c.address = ?", contract).
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
func (s *Storage) NFTs(owner *string, collectionID *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	// Apply explicit query filters - the token owner and collection ID.
	query := api.NFT{}
	if owner != nil {
		query.Owner = *owner
	}
	if collectionID != nil {
		query.Collection = *collectionID
	}
	db := s.db.Where(query)

	// Add the rarity threshold condition.
	if rarityMin != nil {
		db = db.Where("rarity >= ?", rarityMin)
	}

	orderClause, err := formatNFTOrderBy(orderBy)
	if err != nil {
		return nil, fmt.Errorf("could not prepare order clause: %w", err)
	}

	db = db.Order(orderClause)

	var nfts []*api.NFT
	err = db.Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil
}

const (
	rarityColumnName       = "rarity"
	creationTimeColumnName = "created_at"
)

func formatNFTOrderBy(clause api.NFTOrder) (string, error) {

	var field string

	switch clause.Field {

	case api.NFTOrderFieldRarity:
		field = rarityColumnName

	case api.NFTOrderFieldCreationTime:
		field = creationTimeColumnName

	// FIXME: Remove when sorting by value becomes possible.
	default:
		return "", errors.New("unsupported sorting option")
	}

	formatted := fmt.Sprintf("%s %s", field, clause.Direction)

	return formatted, nil
}
