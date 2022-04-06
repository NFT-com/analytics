package storage

import (
	"fmt"

	"github.com/NFT-com/graph-api/events/models/events"
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

	// NOTE: This function creates a query with a limit of `batchSize + 1` to avoid unnecessary queries.
	// See the comment for the `Transfers` query creation for more details.
	limit := s.batchSize + 1
	db, err := s.createQuery(query, token, limit)
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

	// Trim the list to correct size, removing the last element.
	mints = mints[:s.batchSize]

	// Create a token to continue the iteration.
	last := mints[len(mints)-1]
	nextToken := createToken(last.BlockNumber, last.EventIndex)

	return mints, nextToken, nil
}
