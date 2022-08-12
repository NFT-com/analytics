package stats

import (
	"errors"
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionVolume returns the total value of all trades for this collection.
func (s *Stats) CollectionVolume(address identifier.Address) ([]datapoint.Currency, error) {

	query := s.db.
		Table("sales").
		Select("SUM(currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address").
		Where("chain_id = ?", address.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", address.Address).
		Group("chain_id, LOWER(currency_address)")

	var volumes []datapoint.Currency
	err := query.Find(&volumes).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection volume: %w", err)
	}

	return volumes, nil
}

// CollectionBatchVolumes returns the list of volumes for each individual collection.
// Volumes are mapped to the lowercased collection contract address.
func (s *Stats) CollectionBatchVolumes(addresses []identifier.Address) (map[identifier.Address][]datapoint.Currency, error) {

	if len(addresses) == 0 {
		return nil, errors.New("id list must be non-empty")
	}

	query := s.db.
		Table("sales").
		Select("SUM(currency_value) AS currency_value, chain_id, LOWER(collection_address) AS collection_address, LOWER(currency_address) AS currency_address").
		Group("chain_id, LOWER(collection_address), LOWER(currency_address)")

	filter := s.createCollectionFilter(addresses)
	query = query.Where(filter)

	var volumes []batchStatResult
	err := query.Find(&volumes).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection volumes: %w", err)
	}

	// Map the volumes to the collection identifier.
	volumeMap := make(map[identifier.Address][]datapoint.Currency, len(volumes))
	for _, volume := range volumes {

		collection := identifier.Address{
			ChainID: volume.ChainID,
			Address: volume.CollectionAddress,
		}

		currency := datapoint.Currency{
			ChainID: volume.ChainID,
			Address: volume.Address,
			Amount:  volume.Amount,
		}

		// If we already have volume data for this collection (for some currencies)
		// just append the data for this currency.
		_, ok := volumeMap[collection]
		if ok {
			volumeMap[collection] = append(volumeMap[collection], currency)
			continue
		}

		// Create the currency list now.
		v := make([]datapoint.Currency, 0)
		v = append(v, currency)
		volumeMap[collection] = v
	}

	return volumeMap, nil
}

// MarketplaceVolume returns the total value of all trades for this marketplace.
func (s *Stats) MarketplaceVolume(addresses []identifier.Address) ([]datapoint.Currency, error) {

	query := s.db.
		Table("sales").
		Select("SUM(currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address").
		Group("chain_id, LOWER(currency_address)")

	filter := s.createMarketplaceFilter(addresses)
	query = query.Where(filter)

	var volumes []datapoint.Currency
	err := query.Find(&volumes).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplace volume: %w", err)
	}

	return volumes, nil
}
