package stats

import (
	"fmt"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionSales returns the total number of sales for a collection.
func (s *Stats) CollectionSales(address identifier.Address) (datapoint.Sale, error) {

	query := s.db.
		Table("sales").
		Select("COUNT(*) AS count").
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address)

	var count datapoint.Sale
	err := query.Take(&count).Error
	if err != nil {
		return datapoint.Sale{}, fmt.Errorf("could not determine sale count for collection: %w", err)
	}

	return count, nil
}

// MarketplaceSales returns the total number of sales for a marketplace.
func (s *Stats) MarketplaceSales(addresses []identifier.Address) (datapoint.Sale, error) {

	query := s.db.
		Table("sales").
		Select("COUNT(*) AS count")

	filter := s.createMarketplaceFilter(addresses)
	query = query.Where(filter)

	var count datapoint.Sale
	err := query.Take(&count).Error
	if err != nil {
		return datapoint.Sale{}, fmt.Errorf("could not determine sale count for marketplace: %w", err)
	}

	return count, nil
}
