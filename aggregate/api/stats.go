package api

import (
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

type Stats interface {

	// Collection statistics - current.
	CollectionVolume(address identifier.Address) (float64, error)
	CollectionMarketCap(address identifier.Address) (float64, error)
	CollectionSales(address identifier.Address) (uint64, error)
	CollectionBatchVolumes(addresses []identifier.Address) (map[identifier.Address]float64, error)
	CollectionBatchMarketCaps(addresses []identifier.Address) (map[identifier.Address]float64, error)
	CollectionPrices(address identifier.Address) (map[identifier.NFT]float64, error)
	CollectionAveragePrices(address identifier.Address) (map[identifier.NFT]float64, error)

	// Collection statistics - history.
	CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	CollectionSalesHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionLowestPriceHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.LowestPrice, error)
	CollectionAverageHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Average, error)
	CollectionSizeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Count, error)

	// Marketplace statistics - current.
	MarketplaceVolume(addresses []identifier.Address) (float64, error)
	MarketplaceMarketCap(addresses []identifier.Address) (float64, error)
	MarketplaceSales(addresses []identifier.Address) (uint64, error)
	MarketplaceUserCount(addresses []identifier.Address) (uint64, error)

	// Marketplace statistics - history.
	MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	MarketplaceSalesHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error)

	// NFT statistics - current.
	NFTPrice(address identifier.NFT) (datapoint.Currency, error)
	NFTAveragePrice(address identifier.NFT) (datapoint.Average, error)

	// NFT statistics - history.
	NFTPriceHistory(address identifier.NFT, from time.Time, to time.Time) ([]datapoint.Price, error)
}
