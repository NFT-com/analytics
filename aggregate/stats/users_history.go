package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// MarketplaceUserCountHistory returns the number of unique active users on the marketplace in the specified date range.
func (s *Stats) MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error) {

	marketplaceFilter := s.createMarketplaceFilter(addresses)

	// Select all fitting sellers on a marketplace.
	sellersQuery := s.db.
		Table("sales").
		Select("LOWER(seller_address) AS acc").
		Where("emitted_at <= date").
		Where(marketplaceFilter)

	// Select all fitting buyers on a marketplace.
	buyersQuery := s.db.
		Table("sales").
		Select("LOWER(buyer_address) AS acc").
		Where("emitted_at <= date").
		Where(marketplaceFilter)

	// Count query calculates the number of unique users for events in a time span.
	// NOTE: SQL union removes duplicates by default.
	query := s.db.
		Table("generate_series(?::timestamp, ?::timestamp, interval '1 day') as date, "+
			"LATERAL (( ? ) UNION (?)) users",
			from.Format(timeFormat),
			to.Format(timeFormat),
			sellersQuery,
			buyersQuery).
		Select("COUNT (users.*) AS count, date").
		Group("date")

	var out []datapoint.Users
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve number of users: %w", err)
	}
	return out, nil
}
