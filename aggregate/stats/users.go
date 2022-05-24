package stats

import (
	"fmt"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// MarketplaceUserCount returns the total number of unique users for a marketplace.
func (s *Stats) MarketplaceUserCount(addresses []identifier.Address) (datapoint.Users, error) {

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

	query := s.db.
		Table("( ? ) UNION ( ? ) users",
			sellersQuery,
			buyersQuery).
		Select("COUNT(users.*) AS count")

	var count datapoint.Users
	err := query.Take(&count).Error
	if err != nil {
		return datapoint.Users{}, fmt.Errorf("could not retrieve user count: %w", err)
	}

	return count, nil
}
