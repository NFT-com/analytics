package api

import (
	"github.com/shopspring/decimal"
)

// TODO: Add batch requests for all stats.
// See https://github.com/NFT-com/analytics/issues/39.

type Stats interface {
	Prices(nftIDs []string) (map[string]decimal.Decimal, error)
	AveragePrice(nftID string) (decimal.Decimal, error)

	CollectionVolumes(collectionIDs []string) (map[string]decimal.Decimal, error)
	CollectionMarketCaps(collectionIDs []string) (map[string]decimal.Decimal, error)
	CollectionSales(collectionID string) (decimal.Decimal, error)

	MarketplaceVolume(marketplaceID string) (decimal.Decimal, error)
	MarketplaceMarketCap(marketplaceID string) (decimal.Decimal, error)
	MarketplaceSales(marketplaceID string) (decimal.Decimal, error)
	MarketplaceUsers(marketplaceID string) (decimal.Decimal, error)
}
