package api

import (
	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []datapoint.Value `json:"data,omitempty"`
}
