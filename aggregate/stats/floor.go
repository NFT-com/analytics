package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

// FIXME: Add comments for exported functions.

// CollectionFloor returns the floor price for the collection in the given interval.
// Floor price is the lowest price for an NFT in that collection on the given point in time.s
func (s *Stats) CollectionFloor(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Floor, error) {

	intervalQuery := s.db.
		Table("sales, LATERAL generate_series(?, ?, INTERVAL '1 day') AS start_date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select([]string{
			"sales.*",
			"start_date",
			"start_date + interval '1 day' AS end_date"}).
		Where("collection_address = ?", collectionAddress)

	seriesQuery := s.db.
		Table("(?) s", intervalQuery).
		Select("s.*").
		Where("s.emitted_at > s.start_date").
		Where("s.emitted_at <= s.end_date")

	query := s.db.
		Table("(?) st", seriesQuery).
		Select("MIN(st.trade_price) AS floor, st.start_date, st.end_date").
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
