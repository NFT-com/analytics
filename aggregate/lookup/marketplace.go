package lookup

import (
	"fmt"

	aggregate "github.com/NFT-com/analytics/aggregate/api"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// Marketplace returns the list of addresses for the specified marketplace.
func (l *Lookup) Marketplace(id string) ([]identifier.Address, error) {

	query := l.db.
		Table("marketplaces m").
		Joins("INNER JOIN networks_marketplaces nm ON nm.marketplace_id = m.id").
		Joins("INNER JOIN networks n ON nm.network_id = n.id").
		Select("n.chain_id, nm.contract_address").
		Where("m.id = ?", id)

	var addresses []networkAddress
	err := query.Find(&addresses).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplace addresses: %w", err)
	}

	if len(addresses) == 0 {
		return nil, aggregate.ErrRecordNotFound
	}

	out := make([]identifier.Address, 0, len(addresses))
	for _, address := range addresses {

		a := identifier.Address{
			ChainID: address.ChainID,
			Address: address.ContractAddress,
		}

		out = append(out, a)
	}

	return out, nil
}
