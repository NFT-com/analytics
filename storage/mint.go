package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/models/events"
)

func (s *Storage) Mints(filter api.Filter) ([]events.Mint, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
