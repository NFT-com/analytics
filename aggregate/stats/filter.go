package stats

import (
	"gorm.io/gorm"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// FIXME: Create a single `createFilter` function with an enum indicating whether to use
// collection or marketplace addresses for filtering.

// createMarketplaceFilter accepts a list of collection or marketplace aaddresses and returns the
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

// createCollectionFilter accepts a list of collection addresses and returns the appropriate `WHERE`
// clause to the SQL query.
func (s Stats) createCollectionFilter(addresses []identifier.Address) *gorm.DB {

	// Return an empty condition.
	if len(addresses) == 0 {
		return s.db
	}

	// Create the first condition.
	cdb := s.db.
		Where("chain_id = ? AND collection_address = ?",
			addresses[0].ChainID,
			addresses[0].Address,
		)

	// Add the remaining conditions using an `OR`.
	for _, address := range addresses[1:] {
		cdb = cdb.Or("chain_id = ? AND collection_address = ?",
			address.ChainID,
			address.Address,
		)
	}

	return cdb
}

// createNFTFilter accepts a list of NFT identifiers and returns the appropriate
// `WHERE` clauses for the SQL query.
func (s *Stats) createNFTFilter(nfts []identifier.NFT) *gorm.DB {

	// Return an empty condition.
	if len(nfts) == 0 {
		return s.db
	}

	nft := nfts[0]

	// Create the first condition.
	filter := s.db.Where("chain_id = ? AND collection_address = ? AND token_id = ?",
		nft.Collection.ChainID,
		nft.Collection.Address,
		nft.TokenID,
	)

	// Add the remaining conditions using an `OR`.
	for _, nft := range nfts[1:] {
		filter = filter.Or("chain_id = ? AND collection_address = ? AND token_id = ?",
			nft.Collection.ChainID,
			nft.Collection.Address,
			nft.TokenID,
		)
	}

	return filter
}
