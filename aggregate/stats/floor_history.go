package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionFloorHistory returns the floor price for the collection in the given interval.
// Floor price is the lowest price for an NFT in that collection on the given point in time.s
func (s *Stats) CollectionFloorHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Floor, error) {

	intervalQuery := s.db.
		Table("sales, LATERAL generate_series(?::timestamp, ?::timestamp, INTERVAL '1 day') AS start_date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select([]string{
			"sales.*",
			"start_date",
			"start_date + interval '1 day' AS end_date"}).
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address)

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
