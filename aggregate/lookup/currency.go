package lookup

import (
	"fmt"

	aggregate "github.com/NFT-com/analytics/aggregate/api"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// Currency returns the ID of the specified currency.
func (l *Lookup) CurrencyID(currency identifier.Currency) (string, error) {

	query := l.db.
		Table("currencies c, networks n").
		Select("c.id").
		Where("n.id = c.network_id").
		Where("n.chain_id = ?", currency.ChainID).
		Where("c.address = ?", currency.Address).
		Limit(1)

	var ids []string
	err := query.Find(&ids).Error
	if err != nil {
		return "", fmt.Errorf("could not retrieve currency ID: %w", err)
	}

	if len(ids) == 0 {
		return "", aggregate.ErrRecordNotFound
	}

	return ids[0], nil
}
