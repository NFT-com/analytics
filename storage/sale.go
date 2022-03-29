package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Sales retrieves all NFT sale events according to the specified filters.
func (s *Storage) Sales(selector events.SaleSelector) ([]events.Sale, error) {

	query := events.Sale{
		Marketplace: selector.Marketplace,
		Transaction: selector.Transaction,
		Seller:      selector.Seller,
		Buyer:       selector.Buyer,
		Price:       selector.Price,
	}

	// Create the database query.
	db := s.db.Where(query)
	db = setTimeFilter(db, selector.TimeSelector)

	var sales []events.Sale
	err := db.Find(&sales).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve sales events: %w", err)
	}

	return sales, nil
}
