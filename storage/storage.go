package storage

import (
	"fmt"

	"gorm.io/gorm"
)

// FIXME: All 'not found' errors should be sentinels.

// Storage reads the event data from the underlying database.
type Storage struct {
	db *gorm.DB

	batchSize uint
}

// New creates a new Storage handler.
func New(db *gorm.DB, opts ...OptionFunc) *Storage {

	cfg := defaultConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	s := Storage{
		db:        db,
		batchSize: cfg.BatchSize,
	}

	fmt.Printf("batch size: %v", s.batchSize)

	return &s
}
