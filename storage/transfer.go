package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/models/events"
)

// Transfers retrieves all NFT transfer events according to the specified filters.
func (s *Storage) Transfers(selector api.TransferSelector) ([]events.Transfer, error) {

	query := events.Transfer{
		Collection:  selector.Collection,
		Transaction: selector.Transaction,
		TokenID:     selector.TokenID,
		From:        selector.From,
		To:          selector.To,
	}

	// Create the database query.
	db := s.db.Where(query)

	// Set start time condition if provided.
	if selector.Start != "" {
		db = db.Where("emitted_at >= ?", selector.Start)
	}
	// Set end time condition if provided.
	if selector.End != "" {
		db = db.Where("emitted_at <= ?", selector.End)
	}

	var transfers []events.Transfer
	err := db.Find(&transfers).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve transfer events: %w", err)
	}

	return transfers, nil
}
