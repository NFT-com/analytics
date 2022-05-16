package storage

import (
	"gorm.io/gorm"
)

// Storage provides the collection and marketplace lookup functionality.
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
