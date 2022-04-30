package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

func (s *Stats) Count(collectionID string, from time.Time, to time.Time) ([]datapoint.Count, error) {

	countQuery := s.db.
		Table("generate_series(?, ?, INTERVAL '1 day') AS date", from.Format(timeFormat), to.Format(timeFormat)).
		Select("(SELECT COUNT(*) FROM mints WHERE collection = ? AND emitted_at <= date) as mints, "+
			"(SELECT COUNT(*) FROM burns WHERE collection = ? AND emitted_at <= date) as burns, "+
			"date",
			collectionID,
			collectionID)

	query := s.db.
		Table("(?) c", countQuery).
		Select("*").
		Where("c.mints > 0"). // Discard counts before the NFT started being minted.
		Order("date DESC")

	var out []datapoint.Count
	err := query.Find(&out).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT count: %w", err)
	}

	return out, nil
}
