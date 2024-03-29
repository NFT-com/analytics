package aggregate

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/graph/aggregate/http"
)

// Price retrieves the price for the specified NFT.
func (c *Client) Price(id string) ([]api.Coin, error) {

	c.log.Debug().Str("id", id).Msg("requesting NFT price")

	path := fmt.Sprintf(fmtNFTPriceEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeCoinStatRequest(id, address)
}

// AveragePrice retrieves the average price for the specified NFT.
func (c *Client) AveragePrice(id string) ([]api.Coin, error) {

	c.log.Debug().Str("id", id).Msg("requesting NFT average price")

	path := fmt.Sprintf(fmtNFTAveragePriceEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeCoinStatRequest(id, address)
}

// CollectionVolumes retrieves the volumes for the specified collections.
func (c *Client) CollectionVolumes(ids []string) (map[string][]api.Coin, error) {

	c.log.Debug().Strs("ids", ids).Msg("requesting collection volumes")

	address := createAddress(c.apiURL, collectionBatchVolumeEndpoint)
	return c.executeCoinStatBatchRequest(ids, address)
}

// CollectionMarketCaps retrieves the market caps for the specified collections.
func (c *Client) CollectionMarketCaps(ids []string) (map[string][]api.Coin, error) {

	c.log.Debug().Strs("ids", ids).Msg("requesting collection market caps")

	address := createAddress(c.apiURL, collectionBatchMarketCapEndpoint)
	return c.executeCoinStatBatchRequest(ids, address)
}

// CollectionSales retrieves the sale count for the specified collection.
func (c *Client) CollectionSales(id string) (uint64, error) {

	c.log.Debug().Str("id", id).Msg("requesting collection sale count")

	path := fmt.Sprintf(fmtCollectionSalesEndpoint, id)
	address := createAddress(c.apiURL, path)

	var res datapoint.Count
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

// CollectionPrices retrieves the prices for NFTs in the specified collection.
func (c *Client) CollectionPrices(id string) (map[string][]api.Coin, error) {

	c.log.Debug().Str("id", id).Msg("requesting prices for a collection")

	path := fmt.Sprintf(fmtCollectionPricesEndpoint, id)
	address := createAddress(c.apiURL, path)

	// Execute the API request.
	var res api.BatchResponse
	err := http.GET(address, &res)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve prices for a collection: %w", err)
	}

	// Create the output.
	out := make(map[string][]api.Coin)
	for _, price := range res.Data {
		out[price.ID] = price.Value
	}

	return out, nil
}

// CollectionAveragePrices retrieves the prices for NFTs in the specified collection.
func (c *Client) CollectionAveragePrices(id string) (map[string][]api.Coin, error) {

	c.log.Debug().Str("id", id).Msg("requesting average prices for a collection")

	path := fmt.Sprintf(fmtCollectionAveragePricesEndpoint, id)
	address := createAddress(c.apiURL, path)

	// Execute the API request.
	var res api.BatchResponse
	err := http.GET(address, &res)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve average prices for a collection: %w", err)
	}

	// Create the output.
	out := make(map[string][]api.Coin)
	for _, price := range res.Data {
		out[price.ID] = price.Value
	}

	return out, nil
}

// MarketplaceVolume retrieves the volume for the specified marketplace.
func (c *Client) MarketplaceVolume(id string) ([]api.Coin, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace volume")

	path := fmt.Sprintf(fmtMarketplaceVolumeEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeCoinStatRequest(id, address)
}

// MarketplaceMarketCap retrieves the market cap for the specified marketplace.
func (c *Client) MarketplaceMarketCap(id string) ([]api.Coin, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace market cap")

	path := fmt.Sprintf(fmtMarketplaceMarketCapEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeCoinStatRequest(id, address)
}

// MarketplaceSales retrieves the sale count for the specified marketplace.
func (c *Client) MarketplaceSales(id string) (uint64, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace sale count")

	path := fmt.Sprintf(fmtMarketplaceSalesEndpoint, id)
	address := createAddress(c.apiURL, path)

	var res datapoint.Count
	err := http.GET(address, &res)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	// Verify that we have the correct record.
	if res.ID != id {
		return 0, fmt.Errorf("unexpected record returned (want: %s, have: %s): %w", id, res.ID, err)
	}

	return res.Value, nil
}

// MarketplaceUsers retrieves the user count for the specified marketplace.
func (c *Client) MarketplaceUsers(id string) (uint64, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace user count")

	path := fmt.Sprintf(fmtMarketplaceUsersEndpoint, id)
	address := createAddress(c.apiURL, path)

	var res datapoint.Count
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
