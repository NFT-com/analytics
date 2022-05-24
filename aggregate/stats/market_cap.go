package stats

import (
	"fmt"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionMarketCap returns the current market cap for the collection.
func (s *Stats) CollectionMarketCap(address identifier.Address) (datapoint.MarketCap, error) {

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, collection_address, token_id ORDER BY emitted_at DESC) AS rank").
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address)

	// Summarize query will return the sum of all of the freshest prices for
	// NFTs in a collection. The query leverages the "latest price" query as a subquery
	// for the prices.
	sumQuery := s.db.
		Table("( ? ) s", latestPriceQuery).
		Select("SUM(trade_price) AS total").
		Where("rank = 1")

	var marketCap datapoint.MarketCap
	err := sumQuery.Take(&marketCap).Error
	if err != nil {
		return datapoint.MarketCap{}, fmt.Errorf("could not retrieve market cap: %w", err)
	}

	return marketCap, nil
}

// MarketplaceMarketCap returns the current market cap for the marketplace.
func (s *Stats) MarketplaceMarketCap(addresses []identifier.Address) (datapoint.MarketCap, error) {

	// Latest price query will return prices per NFT ranked by freshness.
	// Prices with the lowest rank (closer to 1) will be the most recent ones.
	latestPriceQuery := s.db.
		Table("sales").
		Select("sales.*, row_number() OVER (PARTITION BY chain_id, collection_address, token_id ORDER BY emitted_at DESC) AS rank")

	filter := s.createMarketplaceFilter(addresses)
	latestPriceQuery = latestPriceQuery.Where(filter)

	// Summarize query will return the sum of all of the freshest prices for
	// NFTs in a collection. The query leverages the "latest price" query as a subquery
	// for the prices.
	sumQuery := s.db.
		Table("( ? ) s", latestPriceQuery).
		Select("SUM(trade_price) AS total").
		Where("rank = 1")

	var marketCap datapoint.MarketCap
	err := sumQuery.Take(&marketCap).Error
	if err != nil {
		return datapoint.MarketCap{}, fmt.Errorf("could not retrieve market cap: %w", err)
	}

	return marketCap, nil
}