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

	// If the number of returned items is smaller or equal to `batchSize`,
	// there is no next page of results.
	haveMore := uint(len(burns)) > s.batchSize
	if !haveMore {
		return burns, "", nil
	}

	// The number of records is larger than `batchSize`, meaning there's
	// at least one more page of results - create a token to continue the
	// iteration.

	// Trim the list to correct size, removing the last element.
	burns = burns[:s.batchSize]

	// Create a token for a subsequent search.
	last := burns[len(burns)-1]
	nextToken := createToken(last.BlockNumber, last.EventIndex)

	return burns, nextToken, nil
}
