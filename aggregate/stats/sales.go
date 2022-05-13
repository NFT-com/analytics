package stats

import (
	"errors"
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

func (s *Stats) Sales(chainID uint, collectionAddress string, marketplaceAddress string, from time.Time, to time.Time) ([]datapoint.Sale, error) {

	// Either collection or marketplace address is required.
	if collectionAddress == "" && marketplaceAddress == "" {
		return nil, errors.New("collection or marketplace address is required")
	}

	countQuery := s.db.
		Table("sales, generate_series(?, ?, interval '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select("COUNT(*) AS count, date").
		Where("chain_id = ?", chainID).
		Where("emitted_at <= date").
		Group("date")

	if collectionAddress != "" {
		countQuery = countQuery.Where("collection_address = ?", collectionAddress)
	}
	if marketplaceAddress != "" {
		countQuery = countQuery.Where("marketplace_address = ?", marketplaceAddress)
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
