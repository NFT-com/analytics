package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionVolume returns the total value of all trades in this collection in the given interval.
// Volume for a point in time is calculated as a sum of all sales made until (and including) that moment.
func (s *Stats) CollectionVolume(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error) {

	// FIXME: Use 'timestamp without time zone' for 'generate_series', slight performance improvement

	// Determine the total value of trades for each point in time.
	sumQuery := s.db.
		Select("SUM(trade_price) AS total, date").
		Table("sales, LATERAL generate_series(?, ?, INTERVAL '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address).
		Where("emitted_at <= date").
		Group("date")

	// Determine the difference from the previous data point.
	seriesQuery := s.db.
		Table("(?) s", sumQuery).
		Select("s.total, s.total - LAG(s.total, 1) OVER (ORDER BY date ASC) AS delta, s.date")

	// Only keep those data points where the volume changed.
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
