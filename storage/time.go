package storage

import (
	"gorm.io/gorm"

	"github.com/NFT-com/events-api/models/events"
)

// setTimeFilter will add the time range condition to the query, if required.
func setTimeFilter(db *gorm.DB, selector events.TimeSelector) *gorm.DB {

	// Set start time condition if provided.
	if selector.Start != "" {
		db = db.Where("emitted_at >= ?", selector.Start)
	}
	// Set end time condition if provided.
	if selector.End != "" {
		db = db.Where("emitted_at <= ?", selector.End)
	}

	return db
}
