package api

import (
	"fmt"
	"time"
)

const (
	// TimeFormat represents the expected format for all inputs for the
	// (from, to] date ranges that the API accepts.
	TimeFormat = "2006-01-02"
)

// Request describes the raw API request with an
// ID - path parameter, and a (from, to] date range specified
// via query parameters.
type Request struct {
	ID   string     `param:"id"`
	From RangeBound `query:"from"`
	To   RangeBound `query:"to"`
}

// BatchRequest describes the API request with a list of IDs sent via POST request as JSON.
type BatchRequest struct {
	IDs []string `json:"ids"`
}

// RangeBound is a thin wrapper around time.Time.
// It implements a custom unmarshaller so that the time value and format
// are immediately verified during the echo `Bind` call, and no manual
// time parsing is needed.
type RangeBound time.Time

// UnmarshalParam is used by the echo framework on request binding.
func (b *RangeBound) UnmarshalParam(param string) error {

	t, err := time.Parse(TimeFormat, param)
	if err != nil {
		return fmt.Errorf("invalid range bound (have: %s): %w", param, err)
	}

	*b = RangeBound(t)
	return nil
}

// Time returns the underlying time.Time for the RangeBound.
func (b RangeBound) Time() time.Time {
	return time.Time(b)
}
