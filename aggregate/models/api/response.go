package api

import (
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
)

// Value represents a generic datatype for some currency-related stat.
type Value struct {
	ID    string             `json:"id,omitempty"`
	Value datapoint.Currency `json:"value,omitempty"`
}

// Values represents the generic datatype for some currency-related stat.
type Values struct {
	ID     string               `json:"id,omitempty"`
	Values []datapoint.Currency `json:"values,omitempty"`
}

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []Values `json:"data,omitempty"`
}
