package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/models/events"
)

func (s *Storage) Transfers(filter api.Filter) ([]events.Transfer, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
