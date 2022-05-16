package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

// CollectionSales returns the number of sales in this collection in the given interval.
// For each data point, the number returned indicates the total number of sales up to (and including) that date.
func (s *Stats) CollectionSales(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Sale, error) {

	countQuery := s.db.
		Table("sales, generate_series(?, ?, interval '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select("COUNT(*) AS count, date").
		Where("chain_id = ?", chainID).
		Where("collection_address = ?", collectionAddress).
		Where("emitted_at <= date").
		Group("date")

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
