package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionVolumeHistory returns the total value of all trades in specified collection in the given interval.
// Volume for a point in time is calculated as a sum of all sales made until (and including) that moment.
func (s *Stats) CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error) {
	return s.volumeHistory(&address, nil, from, to)
}

// MarketplaceVolumeHistory returns the total value of all trades in specified marketplace in the given interval.
// Volume for a point in time is calculated as a sum of all sales made until (and including) that moment.
func (s *Stats) MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error) {
	return s.volumeHistory(nil, addresses, from, to)
}

func (s *Stats) volumeHistory(collectionAddress *identifier.Address, marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error) {

	// Determine the total value of trades for each point in time.
	sumQuery := s.db.
		Select("SUM(trade_price) AS total, date").
		Table("sales, LATERAL generate_series(?::timestamp, ?::timestamp, INTERVAL '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Where("emitted_at <= date").
		Group("date")

	// Set collection filter if needed.
	if collectionAddress != nil {
		list := []identifier.Address{
			*collectionAddress,
		}
		collectionFilter := s.createCollectionFilter(list)
		sumQuery.Where(collectionFilter)
	}

	// Set marketplace filter if needed.
	if len(marketplaceAddresses) > 0 {
		marketplaceFilter := s.createMarketplaceFilter(marketplaceAddresses)
		sumQuery.Where(marketplaceFilter)
	}

	// Determine the difference from the previous data point.
	seriesQuery := s.db.
		Table("(?) s", sumQuery).
		Select("s.total, s.total - LAG(s.total, 1) OVER (ORDER BY date ASC) AS delta, s.date")

	// Only keep those data points where the volume changed.
	query := s.db.
		Table("(?) st", seriesQuery).
		Select("st.total, st.date").
		Where("st.delta != 0").Or("st.delta IS NULL").
		Order("date DESC")

	var out []datapoint.Volume
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve volume info: %w", err)
	}

	return out, nil
}
