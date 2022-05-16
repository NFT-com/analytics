package api

import (
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

type Stats interface {
	// FIXME: These are all Collection-only queries at the moment. Soon they may be collection+marketplace again.
	// Collection statistics.
	CollectionVolume(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Volume, error)
	CollectionMarketCap(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	CollectionSales(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionFloor(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Floor, error)
	CollectionAverage(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Average, error)
	CollectionCount(chainID uint, collectionAddress string, from time.Time, to time.Time) ([]datapoint.Count, error)

	// NFT statistics.
	NFTPrice(chainID uint, collectionAddress string, tokenID string, from time.Time, to time.Time) ([]datapoint.Price, error)
	NFTAveragePrice(chainID uint, collectionAddress string, tokenID string) (datapoint.Average, error)

	// Marketplace statistics.
	// MarketplaceUsers(chainID uint, marketplaceAddress string, from time.Time, to time.Time) ([]datapoint.Users, error)
}
