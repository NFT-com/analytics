package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionSales returns the number of sales in this collection in the given interval.
// For each data point, the number returned indicates the total number of sales up to (and including) that date.
func (s *Stats) CollectionSales(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error) {
	return s.sales(&address, nil, from, to)
}

// MarketplaceSales returns the number of sales on this marketplace in the given interval.
// For each data point, the number returned indicates the total number of sales up to (and including) that date.
func (s *Stats) MarketplaceSales(marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error) {
	return s.sales(nil, marketplaceAddresses, from, to)
}

// sales function creates and executes the query to retrieve number of sales fitting a criteria. Typically either collection address or
// marketplace addresses are provided.
func (s *Stats) sales(collectionAddress *identifier.Address, marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error) {

	countQuery := s.db.
		Table("sales, generate_series(?, ?, interval '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select("COUNT(*) AS count, date").
		Where("emitted_at <= date").
		Group("date")

	// Set collection filter if needed.
	if collectionAddress != nil {
		collectionFilter := s.createCollectionFilter(*collectionAddress)
		countQuery.Where(collectionFilter)
	}

	// Set marketplace filter if needed.
	if len(marketplaceAddresses) > 0 {
		marketplaceFilter := s.createMarketplaceFilter(marketplaceAddresses)
		countQuery.Where(marketplaceFilter)
	}

	// Delta query calculates the change since the previous data point.
	deltaQuery := s.db.
		Table("(?) c", countQuery).
		Select("count, count - LAG(count, 1) OVER (ORDER BY date ASC) AS delta, date")

	// Filter query selects only those data points where the metric changed.
	query := s.db.
		Table("(?) ct", deltaQuery).
		Select("ct.count, ct.date").
		Where("ct.delta != 0").
		Or("ct.delta IS NULL").
		Order("date DESC")

	var out []datapoint.Sale
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve number of sales: %w", err)
	}

	return out, nil
}
