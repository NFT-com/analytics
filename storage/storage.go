package storage

import (
	"gorm.io/gorm"
)

// Storage reads the event data from the underlying database.
type Storage struct {
	db *gorm.DB
}

// New creates a new Storage handler.
func New(db *gorm.DB) *Storage {

	s := Storage{
		db: db,
	}

	return &s
}
