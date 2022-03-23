package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/models/events"
)

// Mints retrieves all NFT mint events according to the specified filters.
func (s *Storage) Mints(filter api.Filter) ([]events.Mint, error) {

	query := events.Mint{
		Chain:      filter.Chain,
		Collection: filter.Collection,
		TokenID:    filter.TokenID,
	}

	// Create the database query.
	db := s.db.Where(query)

	// Set start time condition if provided.
	if filter.Start != "" {
		db = db.Where("timestamp > ?", filter.Start)
	}
	// Set end time condition if provided.
	if filter.End != "" {
		db = db.Where("timestamp < ?", filter.End)
	}

	var mints []events.Mint
	err := db.Find(&mints).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve mint events: %w", err)
	}

	return mints, nil
}
