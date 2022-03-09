package storage

import (
	"gorm.io/gorm"
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

// FIXME: Improve error handling everywhere.
