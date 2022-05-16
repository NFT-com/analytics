package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	aggregate "github.com/NFT-com/graph-api/aggregate/api"
)

// CollectionAddress returns the chain ID and contract address for a collection
func (s *Storage) CollectionAddress(id string) (uint, string, error) {

	var address collectionAddress
	err := s.db.
		Table("collections c, networks n").
		Select("n.chain_id, c.contract_address").
		Where("c.id = ?", id).
		Where("n.id = c.network_id").
		First(&address).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, "", aggregate.ErrRecordNotFound
	}
	if err != nil {
		return 0, "", fmt.Errorf("could not retrieve collection address: %w", err)
	}

	return address.ChainID, address.Address, nil
}
