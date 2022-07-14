package api

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

type Stats interface {

	// Collection statistics - current.
	CollectionVolume(address identifier.Address) (decimal.Decimal, error)
	CollectionMarketCap(address identifier.Address) (decimal.Decimal, error)
	CollectionSales(address identifier.Address) (uint64, error)
	CollectionBatchVolumes(addresses []identifier.Address) (map[identifier.Address]decimal.Decimal, error)
	CollectionBatchMarketCaps(addresses []identifier.Address) (map[identifier.Address]decimal.Decimal, error)

	// Collection statistics - history.
	CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	CollectionSalesHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionFloorHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Floor, error)
	CollectionAverageHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Average, error)
	CollectionSizeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Count, error)

	// Marketplace statistics - current.
	MarketplaceVolume(addresses []identifier.Address) (decimal.Decimal, error)
	MarketplaceMarketCap(addresses []identifier.Address) (decimal.Decimal, error)
	MarketplaceSales(addresses []identifier.Address) (uint64, error)
	MarketplaceUserCount(addresses []identifier.Address) (uint64, error)

	// Marketplace statistics - history.
	MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	MarketplaceSalesHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error)

	// NFT statistics - current.
	NFTPrice(address identifier.NFT) (decimal.Decimal, error)
	NFTBatchPrices(addresses []identifier.NFT) (map[identifier.NFT]decimal.Decimal, error)

	// NFT statistics - history.
	NFTPriceHistory(address identifier.NFT, from time.Time, to time.Time) ([]datapoint.Price, error)
	NFTAveragePriceHistory(address identifier.NFT) (datapoint.Average, error)
}
