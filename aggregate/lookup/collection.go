package lookup

import (
	"fmt"

	aggregate "github.com/NFT-com/analytics/aggregate/api"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
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

// Collections returns a set of collection addresses, mapped by their ID.
func (l *Lookup) Collections(ids []string) (map[string]identifier.Address, error) {

	query := l.db.
		Table("collections c, networks n").
		Select("c.id, n.chain_id, c.contract_address").
		Where("n.id = c.network_id").
		Where("c.id IN ?", ids)

	var collections []networkAddress
	err := query.Find(&collections).Error
	if err != nil {
		return nil, fmt.Errorf("could not lookup collection addresses: %w", err)
	}

	// Translate the list of collections to a map.
	addresses := make(map[string]identifier.Address)

	for _, collection := range collections {

		address := identifier.Address{
			ChainID: collection.ChainID,
			Address: collection.ContractAddress,
		}

		addresses[collection.ID] = address
	}

	return addresses, nil
}
