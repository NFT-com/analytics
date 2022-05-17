package storage

import (
	"fmt"

	aggregate "github.com/NFT-com/graph-api/aggregate/api"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// Collection returns the chain ID and contract address for a collection
func (s *Storage) Collection(id string) (identifier.Address, error) {

	// FIXME: Check - using `First` uses wrong table name?

	var address []collectionAddress
	err := s.db.
		Table("collections c, networks n").
		Select("n.chain_id, c.contract_address").
		Where("c.id = ?", id).
		Where("n.id = c.network_id").
		Limit(1).
		Find(&address).Error
	if err != nil {
		return identifier.Address{}, fmt.Errorf("could not retrieve collection address: %w", err)
	}

	if len(address) == 0 {
		return identifier.Address{}, aggregate.ErrRecordNotFound
	}

	out := identifier.Address{
		ChainID: address[0].ChainID,
		Address: address[0].Address,
	}

	return out, nil
}
