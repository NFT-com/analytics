package stats

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionMarketCap returns the current market cap for the collection.
func (s *Stats) CollectionMarketCap(address identifier.Address) (float64, error) {

	query := s.db.Raw(
		`WITH RECURSIVE cte AS (
			(
				SELECT token_id, trade_price
				FROM sales
				WHERE chain_id = ? and LOWER(collection_address) = LOWER(?)
				ORDER BY token_id, emitted_at DESC
				LIMIT 1
			)
			UNION ALL
			SELECT s.*
			FROM cte c,
			LATERAL (
				SELECT token_id, trade_price
				FROM sales
				WHERE token_id > c.token_id AND chain_id = ? and LOWER(collection_address) = LOWER(?)
				ORDER BY token_id, emitted_at DESC
				LIMIT 1
			) s
		)
		SELECT SUM(trade_price) AS total
		FROM cte`,
		address.ChainID,
		address.Address,
		address.ChainID,
		address.Address,
	)

	var marketCap datapoint.MarketCap
	err := query.Scan(&marketCap).Error
	if err != nil {
		return 0, fmt.Errorf("could not retrieve market cap: %w", err)
	}

	return marketCap.Total, nil
}

// CollectionMarketCaps returns the current market cap for the list of collections.
func (s *Stats) CollectionBatchMarketCaps(addresses []identifier.Address) (map[identifier.Address]float64, error) {

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
		Select("SUM(trade_price) AS total, chain_id, LOWER(collection_address) AS collection_address").
		Where("c.rank = 1").
		Group("chain_id, LOWER(collection_address)")

	var caps []batchStatResult
	err := sumQuery.Find(&caps).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve market caps: %w", err)
	}

	// Transform the list of market caps to a map.
	capMap := make(map[identifier.Address]float64, len(caps))
	for _, cap := range caps {

		collection := identifier.Address{
			ChainID: cap.ChainID,
			Address: cap.CollectionAddress,
		}

		capMap[collection] = cap.Total
	}

	return capMap, nil
}

// MarketplaceMarketCap returns the current market cap for the marketplace.
func (s *Stats) MarketplaceMarketCap(addresses []identifier.Address) (float64, error) {

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
		Select("SUM(trade_price) AS total").
		Where("rank = 1")

	var marketCap datapoint.MarketCap
	err := sumQuery.Take(&marketCap).Error
	if err != nil {
		return 0, fmt.Errorf("could not retrieve market cap: %w", err)
	}

	return marketCap.Total, nil
}
