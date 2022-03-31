package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Burns retrieves NFT burn events according to the specified filters.
// Number of events returned is limited by the `batchSize` `Storage` parameter.
// If the number of events for the specified criteria is greater than `batchSize`,
// a token is provided along with the list of events. This token should be provided
// when retrieving the next batch of records.
func (s *Storage) Burns(selector events.BurnSelector, token string) ([]events.Burn, string, error) {

	// Initialize the query variable.
	query := events.Burn{
		Collection:  selector.Collection,
		TokenID:     selector.TokenID,
		Transaction: selector.Transaction,
	}

	// Create the database query.
	db, err := s.createQuery(query, token)
	if err != nil {
		return nil, "", fmt.Errorf("could not create query: %w", err)
	}
	db = setTimeFilter(db, selector.TimeSelector)
	db = setBlockRangeFilter(db, selector.BlockSelector)

	// Retrieve the list of events.
	var burns []events.Burn
	err = db.Find(&burns).Error
	if err != nil {
		return nil, "", fmt.Errorf("could not retrieve burn events: %w", err)
	}

	if len(burns) == 0 {
		return burns, "", nil
	}

	// Create a token for a subsequent search.
	lastID := burns[len(burns)-1].ID
	nextToken := createToken(lastID)

	return burns, nextToken, nil
}
