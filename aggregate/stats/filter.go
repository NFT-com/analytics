package stats

import (
	"gorm.io/gorm"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

const (
	FILTER_COLLECTION = iota + 1
	FILTER_MARKETPLACE
)

func (s *Stats) createCollectionFilter(addresses []identifier.Address) *gorm.DB {
	return s.createAddressFilter(addresses, FILTER_COLLECTION)
}

func (s *Stats) createMarketplaceFilter(addresses []identifier.Address) *gorm.DB {
	return s.createAddressFilter(addresses, FILTER_MARKETPLACE)
}

// createAddressFilter accepts a list of addresses and returns the appropriate `WHERE` clauses
// for the SQL query.
func (s *Stats) createAddressFilter(addresses []identifier.Address, filterType int) *gorm.DB {

	// Return an empty condition if we have no addresses.
	if len(addresses) == 0 {
		return s.db
	}

	var condition string

	// Set the SQL condition to use.
	switch filterType {
	case FILTER_COLLECTION:
		condition = "chain_id = ? AND collection_address = ?"
	case FILTER_MARKETPLACE:
		condition = "chain_id = ? AND marketplace_address = ?"

	// Invalid filter value, just return an empty condition.
	default:
		return s.db
	}

	// Create the first condition.
	filter := s.db.Where(condition, addresses[0].ChainID, addresses[0].Address)

	// Add any remaining conditions with an `OR`.
	for _, address := range addresses[1:] {
		filter = filter.Or(condition, address.ChainID, address.Address)
	}

	return filter
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
