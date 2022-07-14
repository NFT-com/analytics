package aggregate

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Prices retrieves the prices for the specified NFTs.
func (c *Client) Prices(ids []string) (map[string]decimal.Decimal, error) {

	c.log.Debug().Strs("ids", ids).Msg("requesting NFT prices")

	address := createAddress(c.apiURL, nftBatchPriceEndpoint)
	return c.executeBatchRequest(ids, address)
}

// AveragePrice retrieves the average price for the specified NFT.
func (c *Client) AveragePrice(id string) (decimal.Decimal, error) {

	c.log.Debug().Str("id", id).Msg("requesting NFT average price")

	path := fmt.Sprintf(fmtNFTAveragePriceEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeRequest(id, address)
}

// CollectionVolumes retrieves the volumes for the specified collections.
func (c *Client) CollectionVolumes(ids []string) (map[string]decimal.Decimal, error) {

	c.log.Debug().Strs("ids", ids).Msg("requesting collection volumes")

	address := createAddress(c.apiURL, collectionBatchVolumeEndpoint)
	return c.executeBatchRequest(ids, address)
}

// CollectionMarketCaps retrieves the market caps for the specified collections.
func (c *Client) CollectionMarketCaps(ids []string) (map[string]decimal.Decimal, error) {

	c.log.Debug().Strs("ids", ids).Msg("requesting collection market caps")

	address := createAddress(c.apiURL, collectionBatchMarketCapEndpoint)
	return c.executeBatchRequest(ids, address)
}

// CollectionSales retrieves the sale count for the specified collection.
func (c *Client) CollectionSales(id string) (decimal.Decimal, error) {

	c.log.Debug().Str("id", id).Msg("requesting collection sale count")

	path := fmt.Sprintf(fmtCollectionSalesEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeRequest(id, address)
}

// MarketplaceVolume retrieves the volume for the specified marketplace.
func (c *Client) MarketplaceVolume(id string) (decimal.Decimal, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace volume")

	path := fmt.Sprintf(fmtMarketplaceVolumeEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeRequest(id, address)
}

// MarketplaceMarketCap retrieves the market cap for the specified marketplace.
func (c *Client) MarketplaceMarketCap(id string) (decimal.Decimal, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace market cap")

	path := fmt.Sprintf(fmtMarketplaceMarketCapEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeRequest(id, address)
}

// MarketplaceSales retrieves the sale count for the specified marketplace.
func (c *Client) MarketplaceSales(id string) (decimal.Decimal, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace sale count")

	path := fmt.Sprintf(fmtMarketplaceSalesEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeRequest(id, address)
}

// MarketplaceUsers retrieves the user count for the specified marketplace.
func (c *Client) MarketplaceUsers(id string) (decimal.Decimal, error) {

	c.log.Debug().Str("id", id).Msg("requesting marketplace user count")

	path := fmt.Sprintf(fmtMarketplaceUsersEndpoint, id)
	address := createAddress(c.apiURL, path)
	return c.executeRequest(id, address)
}
