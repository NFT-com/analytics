package lookup

import (
	"gorm.io/gorm"
)

// Lookup provides the collection, marketplace and NFT lookup functionality.
type Lookup struct {
	db *gorm.DB
}

// New creates a new Storage handler.
func New(db *gorm.DB) *Lookup {

	l := Lookup{
		db: db,
	}

	return &l
}
