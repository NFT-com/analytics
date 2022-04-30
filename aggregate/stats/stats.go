package stats

import (
	"gorm.io/gorm"
)

const (
	timeFormat = "2006-01-02"
)

// Stats provides NFT statistics based on the Events database.
type Stats struct {
	db *gorm.DB
}

// New creates a new stats handler.
func New(db *gorm.DB) *Stats {

	h := Stats{
		db: db,
	}

	return &h
}
