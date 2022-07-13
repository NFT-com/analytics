package api

// TODO: Add batch requests for all stats.
// See https://github.com/NFT-com/analytics/issues/39.

type Stats interface {
	Prices(nftIDs []string) (map[string]float64, error)
	AveragePrices(nftIDs []string) (map[string]float64, error)

	CollectionVolumes(collectionIDs []string) (map[string]float64, error)
	CollectionMarketCaps(collectionIDs []string) (map[string]float64, error)
	CollectionSales(collectionID string) (float64, error)

	MarketplaceVolume(marketplaceID string) (float64, error)
	MarketplaceMarketCap(marketplaceID string) (float64, error)
	MarketplaceSales(marketplaceID string) (float64, error)
	MarketplaceUsers(marketplaceID string) (float64, error)
}
