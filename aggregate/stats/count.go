package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

const (
	zeroAddress = "0x0000000000000000000000000000000000000000"
)

// FIXME: Make this work now after mints and burns are removed.
func (s *Stats) CollectionCount(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Count, error) {

	mintsQuery := s.db.
		Table("transfers").
		Select("COUNT(*)").
		Where("sender_address = ?", zeroAddress).
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address).
		Where("emitted_at <= date")

	burnsQuery := s.db.
		Table("transfers").
		Select("COUNT(*)").
		Where("receiver_address = ?", zeroAddress).
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address).
		Where("emitted_at <= date")

	countQuery := s.db.
		Table("generate_series(?, ?, INTERVAL '1 day') AS date",
			from.Format(timeFormat),
			to.Format(timeFormat),
		).
		Select("( ? ) AS mints, ( ? ) AS burns, date", mintsQuery, burnsQuery)

	query := s.db.
		Table("( ? ) c", countQuery).
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
