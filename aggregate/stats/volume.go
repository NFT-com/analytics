package stats

import (
	"fmt"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// FIXME: These might be good candidates to support batch requests.

// CollectionVolume returns the total value of all trades for this collection.
func (s *Stats) CollectionVolume(address identifier.Address) (datapoint.Volume, error) {

	query := s.db.
		Table("sales").
		Select("SUM(trade_price) AS total").
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address)

	var volume datapoint.Volume
	err := query.Take(&volume).Error
	if err != nil {
		return datapoint.Volume{}, fmt.Errorf("could not retrieve collection volume: %w", err)
	}

	return volume, nil
}

// MarketplaceVolume returns the total value of all trades for this marketplace.
func (s *Stats) MarketplaceVolume(addresses []identifier.Address) (datapoint.Volume, error) {

	query := s.db.
		Table("sales").
		Select("SUM(trade_price) AS total")

	filter := s.createMarketplaceFilter(addresses)
	query = query.Where(filter)

	var volume datapoint.Volume
	err := query.Take(&volume).Error
	if err != nil {
		return datapoint.Volume{}, fmt.Errorf("could not retrieve marketplace volume: %w", err)
	}

	return volume, nil
}
