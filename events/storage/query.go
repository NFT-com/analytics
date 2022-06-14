package storage

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/NFT-com/analytics/events/models/selectors"
)

type conditionSetFunc func(*gorm.DB) *gorm.DB

// createQuery creates an appropriate event lookup query.
func (s *Storage) createQuery(query interface{}, token string, conditions ...conditionSetFunc) (*gorm.DB, error) {

	db := s.db.
		Where(query).
		Order("block_number DESC").
		Order("event_index DESC")

	// If there's a token provided, unpack it and use it
	// to determine the offset.
	if token != "" {
		height, eventIndex, err := unpackToken(token)
		if err != nil {
			return nil, fmt.Errorf("could not unpack event ID for query offset: %w", err)
		}

		// If this is a continued iteration, request earlier events in the same block
		// and earlier blocks.
		db = db.Where(
			s.db.Where("block_number < ?", height),
		).Or(
			s.db.Where("block_number = ?", height).Where("event_index < ?", eventIndex),
		)
	}

	// Apply conditions provided.
	for _, c := range conditions {
		db = c(db)
	}

	return db, nil
}

// withLimit returns the condition setter that limits the number of results.
func withLimit(limit uint) conditionSetFunc {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Limit(int(limit))
		return db
	}
}

// withTimestampRange returns the condition setter that adds the time range condition
// to the query, if provided.
func withTimestampRange(selector selectors.TimestampRange) conditionSetFunc {
	return func(db *gorm.DB) *gorm.DB {

		// Set start time condition.
		start := time.Time(selector.StartTimestamp)
		if !start.IsZero() {
			db = db.Where("emitted_at >= ?", start.Format(selectors.TimeLayout))
		}

		// Set end time condition.
		end := time.Time(selector.EndTimestamp)
		if !end.IsZero() {
			db = db.Where("emitted_at <= ?", end.Format(selectors.TimeLayout))
		}

		return db
	}
}

// withHeightRange returns the condition setter that adds the height range condition
// to the query, if provided.
func withHeightRange(selector selectors.HeightRange) conditionSetFunc {
	return func(db *gorm.DB) *gorm.DB {

		// Set start height condition.
		if selector.StartHeight != "" {
			db = db.Where("block_number >= ?", selector.StartHeight)
		}

		// Set end height condition.
		if selector.EndHeight != "" {
			db = db.Where("block_number <= ?", selector.EndHeight)
		}

		return db
	}
}

// withPriceRange returns the condition setter that addes the price range condition
// to the query, if provided.
func withPriceRange(selector selectors.PriceRange) conditionSetFunc {
	return func(db *gorm.DB) *gorm.DB {

		// Set the start price condition.
		if selector.StartPrice != 0 {
			db = db.Where("trade_price >= ?", selector.StartPrice)
		}

		// Set end price condition.
		if selector.EndPrice != 0 {
			db = db.Where("trade_price <= ?", selector.EndPrice)
		}

		return db
	}
}
