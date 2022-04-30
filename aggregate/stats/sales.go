package stats

import (
	"errors"
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

func (s *Stats) Sales(collectionID string, marketplaceID string, from time.Time, to time.Time) ([]datapoint.Sale, error) {

	// Either collection or marketplace ID is required.
	if collectionID == "" && marketplaceID == "" {
		return nil, errors.New("collection or marketplace ID is required")
	}

	countQuery := s.db.
		Table("sales_collections, generate_series(?, ?, interval '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select("COUNT(*) AS count, date").
		Where("emitted_at <= date").
		Group("date")

	if collectionID != "" {
		countQuery = countQuery.Where("collection = ?", collectionID)
	}
	if marketplaceID != "" {
		countQuery = countQuery.Where("marketplace = ?", marketplaceID)
	}

	deltaQuery := s.db.
		Table("(?) c", countQuery).
		Select("count, count - LAG(count, 1) OVER (ORDER BY date ASC) AS delta, date")

	query := s.db.
		Table("(?) ct", deltaQuery).
		Select("ct.count, ct.date").
		Where("ct.delta != 0").
		Or("ct.delta IS NULL").
		Order("date DESC")

	var out []datapoint.Sale
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve number of sales: %w", err)
	}

	return out, nil
}
