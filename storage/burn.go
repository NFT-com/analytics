package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Burns retrieves all NFT burn events according to the specified filters.
func (s *Storage) Burns(selector events.BurnSelector) ([]events.Burn, error) {

	query := events.Burn{
		Collection:  selector.Collection,
		TokenID:     selector.TokenID,
		Transaction: selector.Transaction,
	}

	// Create the database query.
	db := s.db.Where(query)
	db = setTimeFilter(db, selector.TimeSelector)

	var burns []events.Burn
	err := db.Find(&burns).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve burn events: %w", err)
	}

	return burns, nil
}
