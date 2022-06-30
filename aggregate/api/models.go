package api

import (
	"time"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// collectionRequest describes the intermediate format for the API request
// for a collection metric, the collection ID translated to a collection address.
type collectionRequest struct {
	address identifier.Address
	from    time.Time
	to      time.Time
}

// marketplaceRequest describes the intermediate format for the API request
// for a marketplace metric, the marketplace ID translated to a list of
// marketplace addresses.
type marketplaceRequest struct {
	addresses []identifier.Address
	from      time.Time
	to        time.Time
}

// nftRequest describes the intermediate format for the API request for an
// NFT metric, the NFT ID translated to an NFT identifier.
type nftRequest struct {
	id   identifier.NFT
	from time.Time
	to   time.Time
}
