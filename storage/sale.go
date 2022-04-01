package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Sales retrieves NFT sale events according to the specified filters.
// Number of events returned is limited by the `batchSize` `Storage` parameter.
// If the number of events for the specified criteria is greater than `batchSize`,
// a token is provided along with the list of events. This token should be provided
// when retrieving the next batch of records.
func (s *Storage) Sales(selector events.SaleSelector, token string) ([]events.Sale, string, error) {

	query := events.Sale{
		Marketplace: selector.Marketplace,
		Transaction: selector.Transaction,
		Seller:      selector.Seller,
		Buyer:       selector.Buyer,
		Price:       selector.Price,
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

	var sales []events.Sale
	err = db.Find(&sales).Error
	if err != nil {
		return nil, "", fmt.Errorf("could not retrieve sales events: %w", err)
	}

	// If the number of returned items is smaller or equal to `batchSize`,
	// there is no next page of results.
	haveMore := uint(len(sales)) > s.batchSize
	if !haveMore {
		return sales, "", nil
	}

	// Trim the list to correct size, removing the last element.
	sales = sales[:s.batchSize]

	// Create a token for a subsequent search.
	last := sales[len(sales)-1]
	nextToken := createToken(last.BlockNumber, last.EventIndex)

	return sales, nextToken, nil
}
