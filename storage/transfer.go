package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// FIXME: Think about retrieving more records than needed to know if there's another page or not.

// Transfers retrieves NFT transfer events according to the specified filters.
// Number of events returned is limited by the `batchSize` `Storage` parameter.
// If the number of events for the specified criteria is greater than `batchSize`,
// a token is provided along with the list of events. This token should be provided
// when retrieving the next batch of records.
func (s *Storage) Transfers(selector events.TransferSelector, token string) ([]events.Transfer, string, error) {

	// Initialize the query variable.
	query := events.Transfer{
		Collection:  selector.Collection,
		Transaction: selector.Transaction,
		TokenID:     selector.TokenID,
		From:        selector.From,
		To:          selector.To,
	}

	// Create the database query.
	db, err := s.createQuery(query, token)
	if err != nil {
		return nil, "", fmt.Errorf("could not create query: %w", err)
	}
	db = setTimeFilter(db, selector.TimeSelector)
	db = setBlockRangeFilter(db, selector.BlockSelector)

	// Retrieve the list of events.
	var transfers []events.Transfer
	err = db.Find(&transfers).Error
	if err != nil {
		return nil, "", fmt.Errorf("could not retrieve transfer events: %w", err)
	}

	if len(transfers) == 0 {
		return transfers, "", nil
	}

	// Create a token for a subsequent search.
	lastID := transfers[len(transfers)-1].ID
	nextToken := createToken(lastID)

	return transfers, nextToken, nil
}
