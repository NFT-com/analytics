package api

import (
	"github.com/NFT-com/analytics/aggregate/models/api"
)

// TODO: Add batch requests for all stats.
// See https://github.com/NFT-com/analytics/issues/39.

type Stats interface {
	Price(nftID string) ([]api.Coin, error)
	AveragePrice(nftID string) ([]api.Coin, error)

	CollectionVolumes(collectionIDs []string) (map[string][]api.Coin, error)
	CollectionMarketCaps(collectionIDs []string) (map[string][]api.Coin, error)
	CollectionSales(collectionID string) (uint64, error)
	CollectionPrices(collectionID string) (map[string][]api.Coin, error)
	CollectionAveragePrices(collectionID string) (map[string][]api.Coin, error)

	MarketplaceVolume(marketplaceID string) ([]api.Coin, error)
	MarketplaceMarketCap(marketplaceID string) ([]api.Coin, error)
	MarketplaceSales(marketplaceID string) (uint64, error)
	MarketplaceUsers(marketplaceID string) (uint64, error)
}
