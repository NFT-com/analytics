package api

import (
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

type Stats interface {

	// Collection statistics - current.
	CollectionVolume(address identifier.Address) (datapoint.Volume, error)
	CollectionMarketCap(address identifier.Address) (datapoint.MarketCap, error)

	// Collection statistics - history.
	CollectionVolumeHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	CollectionMarketCapHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	CollectionSalesHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	CollectionFloorHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Floor, error)
	CollectionAverageHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Average, error)
	CollectionCountHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.Count, error)

	// Marketplace statistics.
	MarketplaceVolumeHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Volume, error)
	MarketplaceMarketCapHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	MarketplaceSalesHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Sale, error)
	MarketplaceUserCountHistory(addresses []identifier.Address, from time.Time, to time.Time) ([]datapoint.Users, error)

	// NFT statistics - current.
	NFTPrice(address identifier.NFT) (datapoint.Price, error)

	// NFT statistics - history.
	NFTPriceHistory(address identifier.NFT, from time.Time, to time.Time) ([]datapoint.Price, error)
	NFTAveragePriceHistory(address identifier.NFT) (datapoint.Average, error)
}
