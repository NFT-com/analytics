package storage

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/NFT-com/graph-api/events/models/events"
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

// withTimeFilter returns the condition setter that adds the time range condition
// to the query, if provided.
func withTimeFilter(selector events.TimeSelector) conditionSetFunc {
	return func(db *gorm.DB) *gorm.DB {

		// Set start time condition.
		start := time.Time(selector.Start)
		if !start.IsZero() {
			db = db.Where("emitted_at >= ?", start.Format(events.TimeLayout))
		}

		// Set end time condition.
		end := time.Time(selector.End)
		if !end.IsZero() {
			db = db.Where("emitted_at <= ?", end.Format(events.TimeLayout))
		}

		return db
	}
}

// withBlockRangeFilter returns the condition setter that adds the block range condition
// to the query, if provided.
func withBlockRangeFilter(selector events.BlockSelector) conditionSetFunc {
	return func(db *gorm.DB) *gorm.DB {

		// Set start block condition.
		if selector.BlockStart != "" {
			db = db.Where("block_number >= ?", selector.BlockStart)
		}
		// Set end block condition.
		if selector.BlockEnd != "" {
			db = db.Where("block_number <= ?", selector.BlockEnd)
		}

		return db
	}
}
