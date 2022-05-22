package api

import (
	"fmt"
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// apiRequest describes the raw API request with an
// ID - path parameter, and a (from, to] date range specified
// via query parameters.
type apiRequest struct {
	ID   string     `param:"id"`
	From rangeBound `query:"from"`
	To   rangeBound `query:"to"`
}

// rangeBound is a thin wrapper around time.Time.
// It implements a custom unmarshaller so that the time value and format
// are immediately verified during the echo `Bind` call. That way
// none of the request handlers have to manually parse time fields
// or check for errors.
type rangeBound time.Time

// UnmarshalParam is used by the echo framework on request binding.
func (b *rangeBound) UnmarshalParam(param string) error {

	t, err := time.Parse(timeFormat, param)
	if err != nil {
		return fmt.Errorf("invalid range bound (have: %s): %w", param, err)
	}

	*b = rangeBound(t)
	return nil
}

func (b rangeBound) time() time.Time {
	return time.Time(b)
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
