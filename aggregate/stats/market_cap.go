package stats

import (
	"errors"
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

// MarketCap returns the market cap for the collection in the given time range.
func (s *Stats) MarketCap(collectionID string, marketplaceID string, from time.Time, to time.Time) ([]datapoint.MarketCap, error) {

	// Either collection or marketplace ID is required.
	if collectionID == "" && marketplaceID == "" {
		return nil, errors.New("collection or marketplace ID is required")
	}

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	// The query has a date threshold to consider only prices up to a date.
	latestPriceQuery := s.db.
		Table("sales_collections").
		Select("sales_collections.*, row_number() OVER (PARTITION BY nft ORDER BY emitted_at DESC) AS rank").
		Where("emitted_at <= d.date")

	if collectionID != "" {
		latestPriceQuery = latestPriceQuery.Where("collection = ?", collectionID)
	}
	if marketplaceID != "" {
		latestPriceQuery = latestPriceQuery.Where("marketplace = ?", marketplaceID)
	}

	// Summarize query will return the sum of all of the freshest prices for
	// NFTs in a collection. The query leverages the "latest price" query as a subquery
	// for the prices. This query is executed via a lateral join to reference the
	// series of date values, so that we have the values calculated for each date
	// in the specified date range.
	sumQuery := s.db.
		Table("(?) s", latestPriceQuery).
		Select("SUM(price) AS total, d.date").
		Where("rank = 1")

	// Market cap query calculates the actual market cap for each date in the
	// specified date range. It also calculates the change from the previous date.
	marketCapQuery := s.db.
		Table("( SELECT generate_series(?, ?, interval '1 day') AS date ) d, LATERAL( ? ) st ",
			from.Format(timeFormat),
			to.Format(timeFormat),
			sumQuery,
		).Select("total, total - LAG(total, 1) OVER (ORDER BY st.date ASC) AS delta, st.date")

	// Finally, this filter query will omit the results of the market cap query
	// where the market cap did not change.
	query := s.db.
		Table("( ? ) mc", marketCapQuery).
		Select("mc.total, mc.date").
		Where("mc.total > 0").
		Where("mc.delta != 0 OR mc.delta IS NULL").
		Order("date DESC")

	var out []datapoint.MarketCap
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve market cap data: %w", err)
	}

	return out, nil
}