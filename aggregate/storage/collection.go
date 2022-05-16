package storage

import (
	"fmt"

	aggregate "github.com/NFT-com/graph-api/aggregate/api"
)

// CollectionAddress returns the chain ID and contract address for a collection
func (s *Storage) CollectionAddress(id string) (uint, string, error) {

	// FIXME: Check - using `First` uses wrong table name

	var address []collectionAddress
	err := s.db.
		Table("collections c, networks n").
		Select("n.chain_id, c.contract_address").
		Where("c.id = ?", id).
		Where("n.id = c.network_id").
		Limit(1).
		Find(&address).Error
	if err != nil {
		return 0, "", fmt.Errorf("could not retrieve collection address: %w", err)
	}

	if len(address) == 0 {
		return 0, "", aggregate.ErrRecordNotFound
	}

	return address[0].ChainID, address[0].Address, nil
}
