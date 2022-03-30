package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/models/events"
)

// Transfers retrieves all NFT transfer events according to the specified filters.
func (s *Storage) Transfers(selector events.TransferSelector) ([]events.Transfer, error) {

	query := events.Transfer{
		Collection:  selector.Collection,
		Transaction: selector.Transaction,
		TokenID:     selector.TokenID,
		From:        selector.From,
		To:          selector.To,
	}

	// Create the database query.
	db := s.createQuery(query)
	db = setTimeFilter(db, selector.TimeSelector)
	db = setBlockRangeFilter(db, selector.BlockSelector)

	var transfers []events.Transfer
	err := db.Find(&transfers).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve transfer events: %w", err)
	}

	return transfers, nil
}
