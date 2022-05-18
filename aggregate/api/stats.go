package api

import (
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

type Stats interface {
	// FIXME: These are all Collection-only queries at the moment. Soon they may be collection+marketplace again.
	// Collection statistics.
	CollectionVolume(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	CollectionMarketCap(address identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	CollectionSales(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionFloor(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Floor, error)
	CollectionAverage(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Average, error)
	CollectionCount(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Count, error)

	// MarketplaceVolume(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	// MarketplaceMarketCap(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	// MarketplaceSales(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	MarketplaceUserCount(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error)

	// NFT statistics.
	NFTPrice(address identifier.NFT, from time.Time, to time.Time) ([]datapoint.Price, error)
	NFTAveragePrice(address identifier.NFT) (datapoint.Average, error)

	// Marketplace statistics.
	// MarketplaceUsers(chainID uint, marketplaceAddress string, from time.Time, to time.Time) ([]datapoint.Users, error)
}
