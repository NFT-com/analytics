package aggregate

import (
	"errors"
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/graph/aggregate/http"
)

// executeCoinStatBatchRequest executes a generic POST request to retrieve stats for a list of IDS.
func (c *Client) executeCoinStatBatchRequest(ids []string, address string) (map[string][]api.Coin, error) {

	// Create the batch price request.
	req := api.BatchRequest{
		IDs: ids,
	}

	// Execute the API request.
	var res api.BatchResponse
	err := http.POST(address, req, &res)
	// Request succeeded but no data returned.
	if err != nil && errors.Is(err, http.ErrNoData) {

		c.log.Debug().Strs("id", ids).Msg("no data received")

		out := make(map[string][]api.Coin)
		return out, nil
	}
	if err != nil {
		return nil, fmt.Errorf("batch request failed: %w", err)
	}

	// Create the output.
	out := make(map[string][]api.Coin)
	for _, p := range res.Data {
		out[p.ID] = p.Value
	}

	return out, nil
}

// executeCoinStatRequest executes a generic GET request to retrieve stat for a single ID.
func (c *Client) executeCoinStatRequest(id string, address string) ([]api.Coin, error) {

	var res api.Value
	err := http.GET(address, &res)
	// Request succeeded but no data returned.
	if err != nil && errors.Is(err, http.ErrNoData) {

		c.log.Debug().Str("id", id).Msg("no data received")
		return []api.Coin{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Verify that we have the correct record.
	if res.ID != id {
		return nil, fmt.Errorf("unexpected record returned (want: %v, have: %v): %w", id, res.ID, err)
	}

	return res.Value, nil
}
