package stats

import (
	"fmt"
	"sort"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CollectionVolumeHistory returns the total value of all trades in specified collection in the given interval.
// Volume for a point in time is calculated as a sum of all sales made until (and including) that moment.
func (s *Stats) CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error) {
	return s.volumeHistory(&address, nil, from, to)
}

// MarketplaceVolumeHistory returns the total value of all trades in specified marketplace in the given interval.
// Volume for a point in time is calculated as a sum of all sales made until (and including) that moment.
func (s *Stats) MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error) {
	return s.volumeHistory(nil, addresses, from, to)
}

func (s *Stats) volumeHistory(collectionAddress *identifier.Address, marketplaceAddresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error) {

	// Determine the total value of trades for each point in time.
	query := s.db.
		Select("SUM(currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address, date").
		Table("sales, LATERAL generate_series(?::timestamp, ?::timestamp, INTERVAL '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Where("emitted_at <= date").
		Group("date, chain_id, LOWER(currency_address)")

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

	var records []datedPriceResult
	err := query.Find(&records).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve volume info: %w", err)
	}

	volumes := createSnapshotList(records)

	return volumes, nil
}

func createSnapshotList(records []datedPriceResult) []datapoint.CoinSnapshot {

	// FIXME: Use a more optimal approach instead of using a map.
	// Since the list is sorted by date, iterate through the records while
	// keeping track of the active date point for which we're processing currencies.
	// After that date is complete, append the created datapoint to the slice, and start
	// composing the next slice element.

	// 1. Create a map where all currencies for a given date are grouped.
	vm := make(map[time.Time][]datapoint.Coin)
	for _, rec := range records {

		date := rec.Date
		currency := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: rec.ChainID,
				Address: rec.Address,
			},
			Amount: rec.Amount,
		}

		_, ok := vm[date]
		if !ok {
			vm[date] = make([]datapoint.Coin, 0, 1)
		}

		vm[date] = append(vm[date], currency)
	}

	// 2. Translate the map to a slice.
	out := make([]datapoint.CoinSnapshot, 0, len(vm))
	for date, volume := range vm {
		date := date
		v := datapoint.CoinSnapshot{
			Date:  date,
			Coins: volume,
		}

		out = append(out, v)
	}

	// 3. Sort the slice.
	sort.Slice(out, func(i, j int) bool {
		return out[i].Date.Before(out[j].Date)
	})

	return out
}
