package aggregate

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/graph/aggregate/http"
)

// executeBatchRequest executes a generic POST request to retrieve stats for a list of IDS.
func (c *Client) executeBatchRequest(ids []string, address string) (map[string]float64, error) {

	// Create the batch price request.
	req := api.BatchRequest{
		IDs: ids,
	}

	// Execute the API request.
	var res api.BatchResponse
	err := http.POST(address, req, &res)
	if err != nil {
		return nil, fmt.Errorf("batch request failed: %w", err)
	}

	// Create the output.
	out := make(map[string]float64)
	for _, price := range res.Data {
		out[price.ID] = price.Value
	}

	return out, nil
}

// executeRequest executes a generic GET request to retrieve stat for a single ID.
func (c *Client) executeRequest(id string, address string) (float64, error) {

	var res datapoint.Value
	err := http.GET(address, &res)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	// Verify that we have the correct record.
	if res.ID != id {
		return 0, fmt.Errorf("unexpected record returned (want: %v, have: %v): %w", id, res.ID, err)
	}

	return res.Value, nil
}
