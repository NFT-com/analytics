package aggregate

import (
	"net/url"
)

const (
	fmtNFTPriceEndpoint        = "/nft/%s/price"
	fmtNFTAveragePriceEndpoint = "/nft/%s/average"

	fmtCollectionSalesEndpoint         = "/collection/%s/sales"
	fmtCollectionPricesEndpoint        = "/collection/%s/prices"
	fmtCollectionAveragePricesEndpoint = "/collection/%s/average_prices"
	collectionBatchVolumeEndpoint      = "/collection/batch/volume"
	collectionBatchMarketCapEndpoint   = "/collection/batch/market_cap"

	fmtMarketplaceVolumeEndpoint    = "/marketplace/%s/volume"
	fmtMarketplaceMarketCapEndpoint = "/marketplace/%s/market_cap"
	fmtMarketplaceSalesEndpoint     = "/marketplace/%s/sales"
	fmtMarketplaceUsersEndpoint     = "/marketplace/%s/users"
)

// createAddress creates the full URL for an HTTP request based on base URL and path.
func createAddress(baseURL url.URL, path string) string {

	out := baseURL
	out.Path = path

	return out.String()
}
