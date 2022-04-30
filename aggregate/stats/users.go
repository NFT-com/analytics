package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

func (s *Stats) MarketplaceUsers(marketplaceID string, from time.Time, to time.Time) ([]datapoint.Users, error) {

	// Count query calculates the number of unique users for events in a time span.
	countQuery := s.db.
		Table("generate_series(?, ?, interval '1 day' ) AS date, "+
			"LATERAL ( (SELECT seller AS acc FROM sales_collections WHERE marketplace = ? AND emitted_at <= date) UNION"+
			"(SELECT buyer AS acc FROM sales_collections WHERE marketplace = ? AND emitted_at <= date ) ) users",
			from.Format(timeFormat),
			to.Format(timeFormat),
			marketplaceID,
			marketplaceID,
		).Select("COUNT(users.*) AS count, date").
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
