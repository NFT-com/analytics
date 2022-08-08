package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// NFTPriceHistory returns the historic prices of an NFT.
func (s *Stats) NFTPriceHistory(nft identifier.NFT, from time.Time, to time.Time) ([]datapoint.Price, error) {

	// NOTE: This query will not return prices for the NFT if there were no sales
	// in the specified date range, unlike all other queries.

	query := s.db.
		Table("sales").
		Select("currency_value, currency_address, emitted_at").
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

	out := make([]datapoint.Price, 0, len(prices))
	for _, p := range prices {

		price := datapoint.Price{
			Currency: datapoint.Currency{
				Amount:  p.Amount,
				Address: p.Address,
			},
			Time: p.Time,
		}

		out = append(out, price)
	}

	return out, nil
}

// NFTAveragePrice returns the all-time average price.
func (s *Stats) NFTAveragePrice(nft identifier.NFT) (datapoint.Average, error) {

	query := s.db.
		Table("sales").
		Select("AVG(trade_price) AS average").
		Where("chain_id = ?", nft.Collection.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", nft.Collection.Address).
		Where("token_id = ?", nft.TokenID)

	var out datapoint.Average
	err := query.Take(&out).Error
	if err != nil {
		return datapoint.Average{}, fmt.Errorf("could not retrieve average price: %w", err)
	}

	return out, nil
}
