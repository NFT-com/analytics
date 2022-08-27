package stats

import (
	"fmt"
	"time"

	aggregate "github.com/NFT-com/analytics/aggregate/api"
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// NFTPriceHistory returns the historic prices of an NFT.
func (s *Stats) NFTPriceHistory(nft identifier.NFT, from time.Time, to time.Time) ([]datapoint.PriceSnapshot, error) {

	// NOTE: This query will not return prices for the NFT if there were no sales
	// in the specified date range, unlike all other queries.

	query := s.db.
		Table("sales").
		Select("currency_value, chain_id, LOWER(currency_address) AS currency_address, emitted_at").
		Where("chain_id = ?", nft.Collection.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", nft.Collection.Address).
		Where("token_id = ?", nft.TokenID).
		Where("emitted_at > ?", from.Format(timeFormat)).
		Where("emitted_at <= ?", to.Format(timeFormat)).
		Order("emitted_at DESC")

	var prices []priceResult
	err := query.Find(&prices).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT prices: %v", err)
	}

	if len(prices) == 0 {
		return nil, aggregate.ErrRecordNotFound
	}

	out := make([]datapoint.PriceSnapshot, 0, len(prices))
	for _, p := range prices {

		price := datapoint.PriceSnapshot{
			Coin: datapoint.Coin{
				Currency: identifier.Currency{
					ChainID: p.ChainID,
					Address: p.Address,
				},
				Value: p.Value,
			},
			Time: p.Time,
		}

		out = append(out, price)
	}

	return out, nil
}

// NFTAveragePrice returns the all-time average price.
func (s *Stats) NFTAveragePrice(nft identifier.NFT) ([]datapoint.Coin, error) {

	query := s.db.
		Table("sales").
		Select("AVG(currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address").
		Where("chain_id = ?", nft.Collection.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", nft.Collection.Address).
		Where("token_id = ?", nft.TokenID).
		Group("chain_id, LOWER(currency_address)")

	var prices []priceResult
	err := query.Find(&prices).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve average price: %w", err)
	}

	if len(prices) == 0 {
		return nil, aggregate.ErrRecordNotFound
	}

	out := make([]datapoint.Coin, 0, len(prices))
	for _, p := range prices {

		price := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: p.ChainID,
				Address: p.Address,
			},
			Value: p.Value,
		}

		out = append(out, price)
	}

	return out, nil
}
