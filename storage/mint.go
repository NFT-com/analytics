package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Mints retrieves all NFT mint events according to the specified filters.
func (s *Storage) Mints(selector events.MintSelector) ([]events.Mint, error) {

	query := events.Mint{
		Collection:  selector.Collection,
		Transaction: selector.Transaction,
		TokenID:     selector.TokenID,
		Owner:       selector.Owner,
	}

	// Create the database query.
	db := s.db.Where(query)

	// Set start time condition if provided.
	if selector.Start != "" {
		db = db.Where("emitted_at >= ?", selector.Start)
	}
	// Set end time condition if provided.
	if selector.End != "" {
		db = db.Where("emitted_at <= ?", selector.End)
	}

	var mints []events.Mint
	err := db.Find(&mints).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve mint events: %w", err)
	}

	return mints, nil
}
