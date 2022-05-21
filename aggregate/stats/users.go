package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// MarketplaceUserCountHistory returns the number of unique active users on the marketplace in the specified date range.
func (s *Stats) MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error) {

	marketplaceFilter := s.createMarketplaceFilter(addresses)

	// Select all fitting sellers on a marketplace.
	sellersQuery := s.db.
		Table("sales").
		Select("seller_address AS acc").
		Where("emitted_at <= date").
		Where(marketplaceFilter)

	// Select all fitting buyers on a marketplace.
	buyersQuery := s.db.
		Table("sales").
		Select("buyer_address AS acc").
		Where("emitted_at <= date").
		Where(marketplaceFilter)

	// Count query calculates the number of unique users for events in a time span.
	// NOTE: SQL union removes duplicates by default.
	countQuery := s.db.
		Table("generate_series(?, ?, interval '1 day') as date, "+
			"LATERAL (( ? ) UNION (?)) users",
			from.Format(timeFormat),
			to.Format(timeFormat),
			sellersQuery,
			buyersQuery).
		Select("COUNT (users.*) AS count, date").
		Group("date")

	// Delta query will show the number of unique users, along with the change from the previous data point.
	deltaQuery := s.db.
		Table("( ? ) uc", countQuery).
		Select("count, count - LAG(count, 1) OVER (ORDER BY date ASC) AS delta, date")

	// Filter query will select only those data points where the number of users changed.
	filter := s.db.
		Table("( ? ) f", deltaQuery).
		Select("count, date").
		Where("delta != 0").
		Or("delta IS NULL").
		Order("date DESC")

	var out []datapoint.Users
	err := filter.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve number of users: %w", err)
	}
	return out, nil
}
