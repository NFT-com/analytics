package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/models/events"
)

// Burns retrieves all NFT burn events according to the specified filters.
func (s *Storage) Burns(filter api.Filter) ([]events.Burn, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
