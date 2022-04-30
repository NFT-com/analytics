package stats

import (
	"errors"
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

func (s *Stats) Volume(collectionID string, marketplaceID string, from time.Time, to time.Time) ([]datapoint.Volume, error) {

	// Either collection or marketplace ID is required.
	if collectionID == "" && marketplaceID == "" {
		return nil, errors.New("collection or marketplace ID is required")
	}

	// FIXME: Use 'timestamp without time zone' for 'generate_series', slight performance improvement

	sumQuery := s.db.
		Select("SUM(price) AS total, date").
		Table("sales_collections, LATERAL generate_series(?, ?, INTERVAL '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Where("emitted_at <= date").
		Group("date")

	if collectionID != "" {
		sumQuery = sumQuery.Where("collection = ?", collectionID)
	}
	if marketplaceID != "" {
		sumQuery = sumQuery.Where("marketplace = ?", marketplaceID)
	}

	seriesQuery := s.db.
		Table("(?) s", sumQuery).
		Select("s.total, s.total - LAG(s.total, 1) OVER (ORDER BY date ASC) AS delta, s.date")

	query := s.db.
		Table("(?) st", seriesQuery).
		Select("st.total, st.date").
		Where("st.delta != 0").Or("st.delta IS NULL").
		Order("date DESC")

	var out []datapoint.Volume
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve volume info: %w", err)
	}

	return out, nil
}
