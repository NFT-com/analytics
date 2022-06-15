package storage

import (
	"fmt"

	"github.com/NFT-com/analytics/events/models/selectors"
	"github.com/NFT-com/indexer/models/events"
)

// Transfers retrieves NFT transfer events according to the specified filters.
// The number of events returned is limited by the `batchSize` `Storage` parameter.
// If the number of events for the specified criteria is greater than `batchSize`,
// a token is provided along with the list of events. This token should be provided
// when retrieving the next batch of records.
func (s *Storage) Transfers(selector selectors.TransferFilter, token string) ([]events.Transfer, string, error) {

	// Create the database query.
	// NOTE: This function creates a query with a limit of `batchSize + 1`.
	// This is done in order to see if there are more records fitting the search
	// criteria after the current batch. If the number of returned records
	// `n <= batchSize`, then there is no next page, and we saved ourselves
	// the cost of doing another database query to do `SELECT COUNT(*) ...`.
	// It is up to the caller to trim the result set to fit the `batchSize`.
	limit := s.batchSize + 1
	filters := []conditionFunc{
		// Explicit matches.
		withUint64Field("chain_id", selector.ChainID),
		withStrField("collection_address", selector.CollectionAddress),
		withStrField("token_id", selector.TokenID),
		withStrField("transaction_hash", selector.TransactionHash),
		withStrField("sender_address", selector.SenderAddress),
		withStrField("receiver_address", selector.ReceiverAddress),

		// Limit and ranges.
		withLimit(limit),
		withTimestampRange(selector.TimestampRange),
		withHeightRange(selector.HeightRange),
	}

	db, err := s.createQuery(token, filters...)
	if err != nil {
		return nil, "", fmt.Errorf("could not create query: %w", err)
	}

	// Retrieve the list of events.
	var transfers []events.Transfer
	err = db.Find(&transfers).Error
	if err != nil {
		return nil, "", fmt.Errorf("could not retrieve transfer events: %w", err)
	}

	// If the number of returned items is smaller or equal to `batchSize`,
	// there is no next page of results.
	lastPage := uint(len(transfers)) <= s.batchSize
	if lastPage {
		return transfers, "", nil
	}

	// Trim the list to correct size, removing the last element.
	transfers = transfers[:s.batchSize]

	// Create a token to continue the iteration.
	last := transfers[len(transfers)-1]
	nextToken := createToken(last.BlockNumber, last.EventIndex)

	return transfers, nextToken, nil
}
