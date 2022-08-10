package api

import (
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

type Stats interface {

	// Collection statistics - current.
	CollectionVolume(address identifier.Address) ([]datapoint.Currency, error)
	CollectionMarketCap(address identifier.Address) ([]datapoint.Currency, error)
	CollectionSales(address identifier.Address) (uint64, error)
	CollectionBatchVolumes(addresses []identifier.Address) (map[identifier.Address][]datapoint.Currency, error)
	CollectionBatchMarketCaps(addresses []identifier.Address) (map[identifier.Address][]datapoint.Currency, error)
	CollectionPrices(address identifier.Address) (map[identifier.NFT]datapoint.Currency, error)
	CollectionAveragePrices(address identifier.Address) (map[identifier.NFT][]datapoint.Currency, error)

	// Collection statistics - history.
	CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CurrencySnapshot, error)
	CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CurrencySnapshot, error)
	CollectionSalesHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionLowestPriceHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.LowestPrice, error)
	CollectionAverageHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CurrencySnapshot, error)
	CollectionSizeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CollectionSize, error)

	// Marketplace statistics - current.
	MarketplaceVolume(addresses []identifier.Address) ([]datapoint.Currency, error)
	MarketplaceMarketCap(addresses []identifier.Address) ([]datapoint.Currency, error)
	MarketplaceSales(addresses []identifier.Address) (uint64, error)
	MarketplaceUserCount(addresses []identifier.Address) (uint64, error)

	// Marketplace statistics - history.
	MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CurrencySnapshot, error)
	MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CurrencySnapshot, error)
	MarketplaceSalesHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error)

	// NFT statistics - current.
	NFTPrice(address identifier.NFT) (datapoint.Currency, error)
	NFTAveragePrice(address identifier.NFT) ([]datapoint.Currency, error)

	// NFT statistics - history.
	NFTPriceHistory(address identifier.NFT, from time.Time, to time.Time) ([]datapoint.Price, error)
}
