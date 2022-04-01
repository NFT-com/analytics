package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Mints retrieves NFT mint events according to the specified filters.
// Number of events returned is limited by the `batchSize` `Storage` parameter.
// If the number of events for the specified criteria is greater than `batchSize`,
// a token is provided along with the list of events. This token should be provided
// when retrieving the next batch of records.
func (s *Storage) Mints(selector events.MintSelector, token string) ([]events.Mint, string, error) {

	// Initialize the query variable.
	query := events.Mint{
		Collection:  selector.Collection,
		Transaction: selector.Transaction,
		TokenID:     selector.TokenID,
		Owner:       selector.Owner,
	}

	// Create the database query.
	db, err := s.createQuery(query, token)
	if err != nil {
		return nil, "", fmt.Errorf("could not create query: %w", err)
	}
	db = setTimeFilter(db, selector.TimeSelector)
	db = setBlockRangeFilter(db, selector.BlockSelector)

	// Retrieve the list of events.
	var mints []events.Mint
	err = db.Find(&mints).Error
	if err != nil {
		return nil, "", fmt.Errorf("could not retrieve mint events: %w", err)
	}

	// If the number of returned items is smaller or equal to `batchSize`,
	// there is no next page of results.
	haveMore := uint(len(mints)) > s.batchSize
	if !haveMore {
		return mints, "", nil
	}

	// The number of records is larger than `batchSize`, meaning there's
	// at least one more page of results - create a token to continue the
	// iteration.

	// Trim the list to correct size, removing the last element.
	mints = mints[:s.batchSize]

	// Create a token for a subsequent search.
	lastID := mints[len(mints)-1].ID
	nextToken := createToken(lastID)

	return mints, nextToken, nil
}
