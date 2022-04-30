package stats

import (
	"errors"
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

// FIXME: Add comments for exported functions.

func (s *Stats) Floor(collectionID string, from time.Time, to time.Time) ([]datapoint.Floor, error) {

	if collectionID == "" {
		return nil, errors.New("collection ID is required")
	}

	intervalQuery := s.db.
		Table("sales_collections, LATERAL generate_series(?, ?, INTERVAL '1 day') AS start_date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select([]string{
			"sales_collections.*",
			"start_date",
			"start_date + interval '1 day' AS end_date"}).
		Where("collection = ?", collectionID)

	seriesQuery := s.db.
		Table("(?) s", intervalQuery).
		Select("s.*").
		Where("s.emitted_at > s.start_date").
		Where("s.emitted_at <= s.end_date")

	query := s.db.
		Table("(?) st", seriesQuery).
		Select("MIN(st.price) AS floor, st.start_date, st.end_date").
		Group("start_date").
		Group("end_date").
		Order("start_date DESC")

	var out []datapoint.Floor
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve floor info: %w", err)
	}

	return out, nil
}
