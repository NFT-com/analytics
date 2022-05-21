package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionMarketCapHistory returns the market cap for the collection in the given time range.
func (s *Stats) CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error) {
	return s.marketCap(&address, nil, from, to)
}

// MarketplaceMarketCapHistory returns the market cap for the marketplace in the given time range.
func (s *Stats) MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error) {
	return s.marketCap(nil, addresses, from, to)
}

func (s *Stats) marketCap(collectionAddress *identifier.Address, marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error) {

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	// The query has a date threshold to consider only prices up to a date.
	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, collection_address, token_id ORDER BY emitted_at DESC) AS rank").
		Where("emitted_at <= d.date")

	// Set collection filter if needed.
	if collectionAddress != nil {
		collectionFilter := s.createCollectionFilter(*collectionAddress)
		latestPriceQuery.Where(collectionFilter)
	}

	// Set marketplace filter if needed.
	if len(marketplaceAddresses) > 0 {
		marketplaceFilter := s.createMarketplaceFilter(marketplaceAddresses)
		latestPriceQuery.Where(marketplaceFilter)
	}

	// Summarize query will return the sum of all of the freshest prices for
	// NFTs in a collection. The query leverages the "latest price" query as a subquery
	// for the prices. This query is executed via a lateral join to reference the
	// series of date values, so that we have the values calculated for each date
	// in the specified date range.
	sumQuery := s.db.
		Table("(?) s", latestPriceQuery).
		Select("SUM(trade_price) AS total, d.date").
		Where("s.rank = 1")

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
