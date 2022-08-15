package api

import (
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
)

// Value represents the generic datatype for some currency-related stat.
type Value struct {
	ID    string           `json:"id,omitempty"`
	Value []datapoint.Coin `json:"values,omitempty"`
}

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []Value `json:"data,omitempty"`
}
