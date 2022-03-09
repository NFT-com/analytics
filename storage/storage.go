package storage

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/NFT-com/indexer-api/models/api"
)

// Storage provides the database interaction functionality, such as retrieving NFTs and Collections
// from the database.
type Storage struct {
	db *gorm.DB
}

// New creates a new Storage handler.
func New(db *gorm.DB) *Storage {

	storage := Storage{
		db: db,
	}

	return &storage
}

// NFT returns a single NFT based on the ID.
func (s *Storage) NFT(id string) (*api.NFT, error) {

	var nft api.NFT
	err := s.db.First(&nft).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve nft: %w", err)
	}

	return &nft, nil
}

func (s *Storage) NFTs() ([]*api.NFT, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}

// Collection will retrieve a single collection based on the ID.
func (s *Storage) Collection(id string) (*api.Collection, error) {

	var collection api.Collection
	err := s.db.First(&collection).Error
	if err != nil {
		// FIXME: err not found is a separate thing
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return &collection, nil
}

func (s *Storage) Collections() ([]*api.Collection, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
