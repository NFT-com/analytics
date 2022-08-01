package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionLowestPriceHistory returns the lowest price for the collection in the given interval.
func (s *Stats) CollectionLowestPriceHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.LowestPrice, error) {

	intervalQuery := s.db.
		Table("sales, LATERAL generate_series(?::timestamp, ?::timestamp, INTERVAL '1 day') AS start_date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select([]string{
			"sales.*",
			"start_date",
			"start_date + interval '1 day' AS end_date"}).
		Where("chain_id = ?", address.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", address.Address)

	seriesQuery := s.db.
		Table("(?) s", intervalQuery).
		Select("s.*").
		Where("s.emitted_at > s.start_date").
		Where("s.emitted_at <= s.end_date")

	query := s.db.
		Table("(?) st", seriesQuery).
		Select("MIN(st.trade_price) AS lowest_price, st.start_date, st.end_date").
		Group("start_date").
		Group("end_date").
		Order("start_date DESC")

	var out []datapoint.LowestPrice
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve lowest price: %w", err)
	}

	return out, nil
}