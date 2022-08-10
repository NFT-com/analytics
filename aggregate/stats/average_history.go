package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionAverageHistory returns the average price for the collection NFT in the given interval.
// Average is calculating by taking the latest price for all NFTs in the collection at the
// given point in time and averaging them.
func (s *Stats) CollectionAverageHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CurrencySnapshot, error) {

	// NOTE: The query in this function is VERY similar to the market cap query,
	// with the difference that it averages the prices instead of adding them.

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	// The query has a date threshold to consider only prices up to a date.
	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY token_id ORDER BY emitted_at DESC) AS rank").
		Where("chain_id = ? ", address.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", address.Address).
		Where("emitted_at <= d.date")

	// Averaging query will return the average of all of the freshest prices for
	// NFTs in a collection. The query leverages the "latest price" query as a subquery
	// for the prices. This query is executed via a lateral join to reference the
	// series of date values, so that we have the values calculated for each date
	// in the specified date range.
	avgQuery := s.db.
		Table("(?) s", latestPriceQuery).
		Select("AVG(currency_value) AS currency_value, LOWER(currency_address) AS currency_address, d.date").
		Where("rank = 1").
		Group("LOWER(currency_address)")

	// Delta query shows the average prices, as well as the difference between the previous
	// data point.
	deltaQuery := s.db.
		Table("( SELECT generate_series(?::timestamp, ?::timestamp, interval '1 day') AS date ) d, LATERAL( ? ) st ",
			from.Format(timeFormat),
			to.Format(timeFormat),
			avgQuery,
		).Select("currency_value, currency_value - LAG(currency_value, 1) OVER (ORDER BY st.date ASC) AS delta, st.date").
		Group("currency_address")

	// Finally, this filter query will omit the results of the average price query
	// where the average did not change.
	query := s.db.
		Table("( ? ) a", deltaQuery).
		Select("a.currency_value, a.currency_address, a.date").
		Where("a.currency_value > 0").
		Where("a.delta != 0 OR a.delta IS NULL").
		Order("date DESC")

	var records []datedPriceResult
	err := query.Find(&records).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve average price for a collection: %w", err)
	}

	averages := createSnapshotList(records)

	return averages, nil
}
