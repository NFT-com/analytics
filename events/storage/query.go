package storage

import (
	"fmt"
	"math/big"
	"time"

	"gorm.io/gorm"

	"github.com/NFT-com/analytics/events/models/selectors"
)

type conditionFunc func(*gorm.DB) *gorm.DB

// createQuery creates an appropriate event lookup query.
func (s *Storage) createQuery(token string, conditions ...conditionFunc) (*gorm.DB, error) {

	db := s.db.
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
		db = db.Where("( block_number < ? OR (block_number = ? AND event_index < ?) )", height, height, eventIndex)
	}

	// Apply conditions provided.
	for _, c := range conditions {
		db = c(db)
	}

	return db, nil
}

// withField returns the condition setter that does a generic match on a field.
func withUint64Field(name string, value uint64) conditionFunc {
	return func(db *gorm.DB) *gorm.DB {

		if value == 0 {
			return db
		}

		query := fmt.Sprintf("%s = ?", name)
		db = db.Where(query, value)

		return db
	}
}

// withStrField returns the condition setter that does a text match on a field.
func withStrField(name string, value string) conditionFunc {
	return func(db *gorm.DB) *gorm.DB {

		if value == "" {
			return db
		}

		query := fmt.Sprintf("%s = ?", name)
		db = db.Where(query, value)

		return db
	}
}

// withStrCIField returns the condition setter that does case-insensitive text match on a field.
func withStrCIField(name string, value string) conditionFunc {
	return func(db *gorm.DB) *gorm.DB {

		if value == "" {
			return db
		}

		query := fmt.Sprintf("LOWER(%s) = LOWER(?)", name)
		db = db.Where(query, value)

		return db
	}
}

// withLimit returns the condition setter that limits the number of results.
func withLimit(limit uint) conditionFunc {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Limit(int(limit))
		return db
	}
}

// withTimestampRange returns the condition setter that adds the time range condition
// to the query, if provided.
func withTimestampRange(selector selectors.TimestampRange) conditionFunc {
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
func withHeightRange(selector selectors.HeightRange) conditionFunc {
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
func withPriceRange(selector selectors.PriceRange) conditionFunc {
	return func(db *gorm.DB) *gorm.DB {

		// Set the start price condition.
		if selector.StartPrice.Cmp(big.NewInt(0)) != 0 {
			db = db.Where("trade_price >= ?", selector.StartPrice.String())
		}

		// Set end price condition.
		if selector.EndPrice.Cmp(big.NewInt(0)) != 0 {
			db = db.Where("trade_price <= ?", selector.EndPrice.String())
		}

		return db
	}
}
