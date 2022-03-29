package storage

import (
	"gorm.io/gorm"

	"github.com/NFT-com/events-api/models/events"
)

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
