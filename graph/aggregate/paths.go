package aggregate

import (
	"net/url"
)

const (
	nftBatchPriceEndpoint        = "/nft/batch/price"
	nftBatchAveragePriceEndpoint = "/nft/batch/average"

	collectionBatchVolumeEndpoint    = "/collection/batch/volume"
	collectionBatchMarketCapEndpoint = "/collection/batch/market_cap"
	fmtCollectionSalesEndpoint       = "/collection/%s/sales"

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
