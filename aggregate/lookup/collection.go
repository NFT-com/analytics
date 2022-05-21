package lookup

import (
	"fmt"

	aggregate "github.com/NFT-com/graph-api/aggregate/api"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// Collection returns the address of the specified collection.
func (l *Lookup) Collection(id string) (identifier.Address, error) {

	// Note: Using `Find` with a limit of 1 instead of `First` because the generated SQL
	// uses the wrong table name otherwise.

	var address []networkAddress
	err := l.db.
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
		Address: address[0].ContractAddress,
	}

	return out, nil
}
