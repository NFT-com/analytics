package api

import (
	"strings"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// lowerNFTID adjusts the NFT identifier by lowercasing the collection address.
func lowerNFTID(id identifier.NFT) identifier.NFT {

	lowID := identifier.NFT{
		Collection: lowerAddress(id.Collection),
		TokenID:    id.TokenID,
	}

	return lowID
}

// lowerAddress adjusts the address identifier by lowercasing the collection address.
// Collection addresses are often lowercased in database queries for case-insensitive matching.
// By doing the same here, we can use the resulting NFT identifier as a key in a map lookup.
func lowerAddress(id identifier.Address) identifier.Address {

	lowAddr := identifier.Address{
		ChainID: id.ChainID,
		Address: strings.ToLower(id.Address),
	}

	return lowAddr
}
