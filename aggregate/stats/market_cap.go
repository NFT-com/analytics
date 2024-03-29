package stats

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionMarketCap returns the current market cap for the collection.
func (s *Stats) CollectionMarketCap(address identifier.Address) ([]datapoint.Coin, error) {

	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, LOWER(collection_address), token_id ORDER BY emitted_at DESC) AS rank").
		Where("chain_id = ?", address.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", address.Address)

	query := s.db.
		Table("( ? ) c", latestPriceQuery).
		Select("SUM(currency_value) AS currency_value, LOWER(currency_address) AS currency_address, chain_id, LOWER(collection_address) AS collection_address").
		Where("c.rank = 1").
		Group("chain_id, LOWER(collection_address), LOWER(currency_address)")

	var marketCap []priceResult
	err := query.Find(&marketCap).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve market cap: %w", err)
	}

	out := convertPricesToCoins(marketCap)

	return out, nil
}

// CollectionBatchMarketCaps returns the map of current market caps for the list of collections.
// Market caps are mapped to the lowercased collection contract address.
func (s *Stats) CollectionBatchMarketCaps(addresses []identifier.Address) (map[identifier.Address][]datapoint.Coin, error) {

	if len(addresses) == 0 {
		return nil, fmt.Errorf("address list must be non-empty")
	}

	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, LOWER(collection_address), token_id ORDER BY emitted_at DESC) AS rank")

	filter := s.createCollectionFilter(addresses)
	latestPriceQuery = latestPriceQuery.Where(filter)

	sumQuery := s.db.
		Table("( ? ) c", latestPriceQuery).
		Select("SUM(currency_value) AS currency_value, LOWER(currency_address) AS currency_address, chain_id, LOWER(collection_address) AS collection_address").
		Where("c.rank = 1").
		Group("chain_id, LOWER(collection_address), LOWER(currency_address)")

	var caps []batchStatResult
	err := sumQuery.Find(&caps).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve market caps: %w", err)
	}

	// Transform the list of market caps to a map.
	capMap := make(map[identifier.Address][]datapoint.Coin, len(caps))
	for _, mcap := range caps {

		collection := identifier.Address{
			ChainID: mcap.ChainID,
			Address: mcap.CollectionAddress,
		}

		currency := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: mcap.ChainID,
				Address: mcap.Address,
			},
			Value: mcap.Value,
		}

		// If we already have market cap for this collection (for some currencies)
		// just append the data for this currency.
		_, ok := capMap[collection]
		if ok {
			capMap[collection] = append(capMap[collection], currency)
			continue
		}

		// Otherwise, initialize the currency slice.
		c := make([]datapoint.Coin, 0)
		c = append(c, currency)
		capMap[collection] = c
	}

	return capMap, nil
}

// MarketplaceMarketCap returns the current market cap for the marketplace.
func (s *Stats) MarketplaceMarketCap(addresses []identifier.Address) ([]datapoint.Coin, error) {

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, LOWER(collection_address), token_id ORDER BY emitted_at DESC) AS rank")

	filter := s.createMarketplaceFilter(addresses)
	latestPriceQuery = latestPriceQuery.Where(filter)

	// Summarize query will return the sum of all of the freshest prices for
	// NFTs in a collection. The query leverages the "latest price" query as a subquery
	// for the prices.
	sumQuery := s.db.
		Table("( ? ) s", latestPriceQuery).
		Select("SUM(currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address").
		Group("chain_id").
		Group("LOWER(currency_address)").
		Where("rank = 1")

	var marketCaps []priceResult
	err := sumQuery.Find(&marketCaps).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve market cap: %w", err)
	}

	out := convertPricesToCoins(marketCaps)

	return out, nil
}
