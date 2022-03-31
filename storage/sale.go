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

	// Create the database query.
	db, err := s.createQuery(query, token)
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

	if len(sales) == 0 {
		return sales, "", nil
	}

	// Create a token for a subsequent search.
	lastID := sales[len(sales)-1].ID
	nextToken := createToken(lastID)

	return sales, nextToken, nil
}
