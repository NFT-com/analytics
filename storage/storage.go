package storage

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/NFT-com/indexer-api/models/api"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {

	storage := Storage{
		db: db,
	}

	return &storage
}

func (s *Storage) NFT(id string) (*api.Nft, error) {

	flat := flatNFT{
		ID: id,
	}
	err := s.db.First(&flat).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	nft := api.Nft{
		ID:      flat.ID,
		TokenID: flat.TokenID,
		Owner:   flat.Owner,
		URI:     flat.URI,
		Rarity:  flat.Rarity,
	}

	return &nft, nil
}

func (s *Storage) NFTs() ([]*api.Nft, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}

func (s *Storage) Collection(id string) (*api.Collection, error) {

	flat := flatCollection{
		ID: id,
	}
	err := s.db.First(&flat).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	collection := api.Collection{
		ID:          flat.ID,
		Name:        flat.Name,
		Description: flat.Description,
		Address:     flat.Address,
	}

	return &collection, nil
}

func (s *Storage) Collections() ([]*api.Collection, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
