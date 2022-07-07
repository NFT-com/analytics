package stats

import (
	"errors"
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionVolume returns the total value of all trades for this collection.
func (s *Stats) CollectionVolume(address identifier.Address) (float64, error) {

	query := s.db.
		Table("sales").
		Select("SUM(trade_price) AS total").
		Where("chain_id = ?", address.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", address.Address)

	var volume datapoint.Volume
	err := query.Take(&volume).Error
	if err != nil {
		return 0, fmt.Errorf("could not retrieve collection volume: %w", err)
	}

	return volume.Total, nil
}

// CollectionBatchVolumes returns the list of volumes for each individual collection.
func (s *Stats) CollectionBatchVolumes(addresses []identifier.Address) (map[identifier.Address]float64, error) {

	if len(addresses) == 0 {
		return nil, errors.New("id list must be non-empty")
	}

	query := s.db.
		Table("sales").
		Select("SUM(trade_price) AS total, chain_id, LOWER(collection_address)").
		Group("chain_id, LOWER(collection_address)")

	filter := s.createCollectionFilter(addresses)
	query = query.Where(filter)

	var volumes []batchStatResult
	err := query.Find(&volumes).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection volumes: %w", err)
	}

	// Map the volumes to the collection identifier.
	volumeMap := make(map[identifier.Address]float64, len(volumes))
	for _, volume := range volumes {

		collection := identifier.Address{
			ChainID: volume.ChainID,
			Address: volume.CollectionAddress,
		}

		volumeMap[collection] = volume.Total
	}

	return volumeMap, nil
}

// MarketplaceVolume returns the total value of all trades for this marketplace.
func (s *Stats) MarketplaceVolume(addresses []identifier.Address) (float64, error) {

	query := s.db.
		Table("sales").
		Select("SUM(trade_price) AS total")

	filter := s.createMarketplaceFilter(addresses)
	query = query.Where(filter)

	var volume datapoint.Volume
	err := query.Take(&volume).Error
	if err != nil {
		return 0, fmt.Errorf("could not retrieve marketplace volume: %w", err)
	}

	return volume.Total, nil
}
