package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionMarketCapHistory returns the market cap for the collection in the given time range.
func (s *Stats) CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error) {
	return s.marketCapHistory(&address, nil, from, to)
}

// MarketplaceMarketCapHistory returns the market cap for the marketplace in the given time range.
func (s *Stats) MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error) {
	return s.marketCapHistory(nil, addresses, from, to)
}

func (s *Stats) marketCapHistory(collectionAddress *identifier.Address, marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error) {

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	// The query has a date threshold to consider only prices up to a date.
	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, LOWER(collection_address), token_id ORDER BY emitted_at DESC) AS rank").
		Where("emitted_at <= d.date")

	// Set collection filter if needed.
	if collectionAddress != nil {

		list := []identifier.Address{
			*collectionAddress,
		}
		collectionFilter := s.createCollectionFilter(list)
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
		Select("SUM(currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address, d.date").
		Group("chain_id, LOWER(currency_address)").
		Where("s.rank = 1")

	// Market cap query calculates the actual market cap for each date in the
	// specified date range.
	marketCapQuery := s.db.
		Table("( SELECT generate_series(?::timestamp, ?::timestamp, interval '1 day') AS date ) d, LATERAL( ? ) st ",
			from.Format(timeFormat),
			to.Format(timeFormat),
			sumQuery,
		).Select("currency_value, chain_id, LOWER(currency_address) AS currency_address, st.date")

	var records []datedPriceResult
	err := marketCapQuery.Find(&records).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve volume info: %w", err)
	}

	mcap := createCoinSnapshotList(records)

	return mcap, nil
}
