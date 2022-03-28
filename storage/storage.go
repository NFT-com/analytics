package storage

import (
	"gorm.io/gorm"
)

// TODO: Refactor the storage package so it's more generic.

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
