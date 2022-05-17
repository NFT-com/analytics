package stats

import (
	"errors"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// FIXME: Make this work now after mints and burns are removed.
func (s *Stats) CollectionCount(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Count, error) {
	return nil, errors.New("TBD: Not implemented")
}
