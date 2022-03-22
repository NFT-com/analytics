package storage

import (
	"fmt"

	"github.com/NFT-com/events-api/api"
	"github.com/NFT-com/events-api/models/events"
)

func (s *Storage) Sales(filter api.Filter) ([]events.Sale, error) {
	return nil, fmt.Errorf("TBD: not implemented")
}
