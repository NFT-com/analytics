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

type flatNFT struct {
	ID      string
	TokenID string
	Owner   string
	URI     string
	Rarity  float64
}

func (s *Storage) NFT(id string) (*api.Nft, error) {

	var flat flatNFT
	err := s.db.Raw("SELECT * FROM nft WHERE ID = ?", id).Scan(&flat).Error
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

	var nfts []*api.Nft
	err := s.db.Debug().Find(&nfts).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve nfts: %w", err)
	}

	return nfts, nil

}

func (s *Storage) Collection(id string) (*api.Collection, error) {

	var collection api.Collection
	err := s.db.Debug().First(&collection, id).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return &collection, nil
}

func (s *Storage) Collections() ([]*api.Collection, error) {

	var collections []*api.Collection
	err := s.db.Debug().Find(&collections).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}
