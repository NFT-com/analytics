package api

import (
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// apiRequest describes the raw API request with an
// ID - path parameter, and a (from, to] date range specified
// via query parameters.
type apiRequest struct {
	ID   string    `param:"id" validate:"uuid4,required"`
	From time.Time `query:"from" validate:"datetime=2006-01-02"`
	To   time.Time `query:"to" validate:"datetime=2006-01-02,gtfield=From"`
}

// FIXME: Fix these fugly comments.

// collectionRequest describes the API request for a collection metric
// in the interim format.
type collectionRequest struct {
	address identifier.Address
	from    time.Time
	to      time.Time
}

// marketplaceRequest describes the API request for a marketplace metric
// in the interim format.
type marketplaceRequest struct {
	addresses []identifier.Address
	from      time.Time
	to        time.Time
}
