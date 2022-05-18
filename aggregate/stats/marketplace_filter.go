package stats

import (
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
	"gorm.io/gorm"
)

// createMarketplaceFilter accepts a list of marketplace addresses and adds the
// appropriate `WHERE` clauses to the SQL query.
func (s *Stats) createMarketplaceFilter(addresses []identifier.Address) *gorm.DB {

	if len(addresses) == 0 {
		return nil
	}

	mdb := s.db.Where("chain_id = ? AND marketplace_address = ?",
		addresses[0].ChainID,
		addresses[0].Address)

	for _, address := range addresses[1:] {
		mdb = mdb.Or("chain_id = ? AND marketplace_address = ?",
			address.ChainID,
			address.Address)
	}

	return mdb
}
