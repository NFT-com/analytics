package api

import (
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

type Stats interface {
	CurrentCollectionStats
	HistoricalCollectionStats

	CurrentMarketplaceStats
	HistoricalMarketplaceStats

	CurrentNFTStats
	HistoricalNFTStats
}

type CurrentCollectionStats interface {
	CollectionVolume(address identifier.Address) ([]datapoint.Coin, error)
	CollectionMarketCap(address identifier.Address) ([]datapoint.Coin, error)
	CollectionSales(address identifier.Address) (uint64, error)
	CollectionBatchVolumes(addresses []identifier.Address) (map[identifier.Address][]datapoint.Coin, error)
	CollectionBatchMarketCaps(addresses []identifier.Address) (map[identifier.Address][]datapoint.Coin, error)
	CollectionPrices(address identifier.Address) (map[identifier.NFT][]datapoint.Coin, error)
	CollectionAveragePrices(address identifier.Address) (map[identifier.NFT][]datapoint.Coin, error)
}

type HistoricalCollectionStats interface {
	CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error)
	CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error)
	CollectionSalesHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionLowestPriceHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.LowestPrice, error)
	CollectionAverageHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error)
	CollectionSizeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.CollectionSize, error)
}

type CurrentMarketplaceStats interface {
	MarketplaceVolume(addresses []identifier.Address) ([]datapoint.Coin, error)
	MarketplaceMarketCap(addresses []identifier.Address) ([]datapoint.Coin, error)
	MarketplaceSales(addresses []identifier.Address) (uint64, error)
	MarketplaceUserCount(addresses []identifier.Address) (uint64, error)
}

type HistoricalMarketplaceStats interface {
	MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error)
	MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.CoinSnapshot, error)
	MarketplaceSalesHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error)
}

type CurrentNFTStats interface {
	NFTPrice(address identifier.NFT) ([]datapoint.Coin, error)
	NFTAveragePrice(address identifier.NFT) ([]datapoint.Coin, error)
}

type HistoricalNFTStats interface {
	NFTPriceHistory(address identifier.NFT, from time.Time, to time.Time) ([]datapoint.PriceSnapshot, error)
}
