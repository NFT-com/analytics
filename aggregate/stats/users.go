package stats

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// MarketplaceUserCount returns the total number of unique users for a marketplace.
func (s *Stats) MarketplaceUserCount(addresses []identifier.Address) (uint64, error) {

	// NOTE: The query below does a bit of a roundabout way of retrieving the user count.
	// It does 'UNION ALL' instead of 'UNION' (which deduplicates results), and does
	// 'SELECT COUNT' in a separate subquery. However, this variant was at least 7x faster.

	marketplaceFilter := s.createMarketplaceFilter(addresses)

	// Select all fitting sellers on a marketplace.
	sellersQuery := s.db.
		Table("sales").
		Select("seller_address AS acc").
		Where(marketplaceFilter)

	// Select all fitting buyers on a marketplace.
	buyersQuery := s.db.
		Table("sales").
		Select("buyer_address AS acc").
		Where(marketplaceFilter)

	// Select all unique users.
	usersQuery := s.db.
		Table("( ( ? ) UNION ALL ( ? )) users",
			sellersQuery,
			buyersQuery).
		Select("DISTINCT users.acc")

	// Select count of unique users.
	query := s.db.
		Table("( ? ) c", usersQuery).
		Select("COUNT(*)")

	var count datapoint.Users
	err := query.Take(&count).Error
	if err != nil {
		return 0, fmt.Errorf("could not retrieve user count: %w", err)
	}

	return count.Count, nil
}
