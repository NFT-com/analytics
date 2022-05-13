package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

// FIXME: Think about the approach for NFT-price - use the per-day method?

// NFTPrice returns the historic prices of an NFT.
func (s *Stats) NFTPrice(chainID uint, collectionAddress string, tokenID string, from time.Time, to time.Time) ([]datapoint.Price, error) {

	// NOTE: This query will not return prices for the NFT if there were no sales
	// in the specified date range, unlike all other queries.

	query := s.db.
		Table("sales").
		Select("price, emitted_at").
		Where("chain_id = ?", chainID).
		Where("collection_address = ?", collectionAddress).
		Where("token_id = ?", tokenID).
		Where("emitted_at > ?", from.Format(timeFormat)).
		Where("emitted_at <= ?", to.Format(timeFormat)).
		Order("emitted_at DESC")

	var out []datapoint.Price
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT prices: %v", err)
	}

	return out, nil
}

// NFTAveragePrice returns the all-time average price.
func (s *Stats) NFTAveragePrice(chainID uint, collectionAddress string, tokenID string) (datapoint.Average, error) {

	query := s.db.
		Table("sales").
		Select("AVG(trade_price) AS average").
		Where("chain_id = ?", chainID).
		Where("collection_address = ?", collectionAddress).
		Where("token_id = ?", tokenID)

	var out datapoint.Average
	err := query.Take(&out).Error
	if err != nil {
		return datapoint.Average{}, fmt.Errorf("could not retrieve average price: %w", err)
	}

	return out, nil
}
