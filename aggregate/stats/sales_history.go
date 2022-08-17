package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionSalesHistory returns the number of sales in this collection in the given interval.
// For each data point, the number returned indicates the total number of sales up to (and including) that date.
func (s *Stats) CollectionSalesHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error) {
	return s.salesHistory(&address, nil, from, to)
}

// MarketplaceSalesHistory returns the number of sales on this marketplace in the given interval.
// For each data point, the number returned indicates the total number of sales up to (and including) that date.
func (s *Stats) MarketplaceSalesHistory(marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error) {
	return s.salesHistory(nil, marketplaceAddresses, from, to)
}

// salesHistory function creates and executes the query to retrieve number of salesHistory fitting a criteria. Typically either collection address or
// marketplace addresses are provided.
func (s *Stats) salesHistory(collectionAddress *identifier.Address, marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error) {

	query := s.db.
		Table("sales, generate_series(?::timestamp, ?::timestamp, interval '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select("COUNT(*) AS count, date").
		Where("emitted_at <= date").
		Group("date")

	// Set collection filter if needed.
	if collectionAddress != nil {
		list := []identifier.Address{
			*collectionAddress,
		}
		collectionFilter := s.createCollectionFilter(list)
		query.Where(collectionFilter)
	}

	// Set marketplace filter if needed.
	if len(marketplaceAddresses) > 0 {
		marketplaceFilter := s.createMarketplaceFilter(marketplaceAddresses)
		query.Where(marketplaceFilter)
	}

	var out []datapoint.Sale
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve number of sales: %w", err)
	}

	return out, nil
}
