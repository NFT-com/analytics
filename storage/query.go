package storage

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/NFT-com/events-api/models/events"
)

// createQuery creates an appropriate event lookup query.
// NOTE: This function creates a query with a limit of `batchSize + 1`.
// This is done in order to see if there are more records fitting the search
// criteria after the current batch. If the number of returned records
// `n <= batchSize`, then there is no next page, and we saved ourselves
// the cost of doing another database query to do `SELECT COUNT(*) ...`.
// It is up to the caller to trim the result set to fit the `batchSize`.
func (s *Storage) createQuery(query interface{}, token string) (*gorm.DB, error) {

	db := s.db.
		Where(query).
		Limit(int(s.batchSize + 1)).
		Order("block DESC").
		Order("event_index DESC")

	// If we don't have a token, we're done.
	if token == "" {
		return db, nil
	}

	// If there's a token provided, unpack it and use it
	// to determine the offset.

	blockNo, eventIndex, err := unpackToken(token)
	if err != nil {
		return nil, fmt.Errorf("could not unpack event ID for query offset: %w", err)
	}

	// If this is a continued iteration, request earlier events in the same block
	// and earlier blocks.
	db = db.Where(
		s.db.Where("block < ?", blockNo),
	).Or(
		s.db.Where("block = ?", blockNo).Where("event_index < ?", eventIndex),
	)

	return db, nil
}

// setTimeFilter will add the time range condition to the query, if provided.
func setTimeFilter(db *gorm.DB, selector events.TimeSelector) *gorm.DB {

	// Set start time condition.
	if selector.Start != "" {
		db = db.Where("emitted_at >= ?", selector.Start)
	}

	// Set end time condition.
	if selector.End != "" {
		db = db.Where("emitted_at <= ?", selector.End)
	}

	return db
}

// setBlockRangeFilter will add the block range condition to the query, if provided.
func setBlockRangeFilter(db *gorm.DB, selector events.BlockSelector) *gorm.DB {

	// Set start block condition.
	if selector.BlockStart != "" {
		db = db.Where("block >= ?", selector.BlockStart)
	}
	// Set end block condition.
	if selector.BlockEnd != "" {
		db = db.Where("block <= ?", selector.BlockEnd)
	}

	return db
}
