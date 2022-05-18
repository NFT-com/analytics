package stats

import (
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
	"gorm.io/gorm"
)

// createMarketplaceFilter accepts a list of marketplace addresses and adds the
// appropriate `WHERE` clauses to the SQL query.
func (s *Stats) createMarketplaceFilter(addresses []identifier.Address) *gorm.DB {

	// Return an empty condition.
	if len(addresses) == 0 {
		return s.db
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

// createCollectionFilter accepts a collection address and adds the appropriate `WHERE` clause to the
// SQL query.
func (s Stats) createCollectionFilter(address identifier.Address) *gorm.DB {

	cdb := s.db.
		Where("chain_id = ?", address.ChainID).
		Where("collection_address = ?", address.Address)

	return cdb
}
