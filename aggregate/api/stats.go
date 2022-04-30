package api

import (
	"time"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
)

type Stats interface {
	Volume(collectionID string, marketplaceID string, from time.Time, to time.Time) ([]datapoint.Volume, error)
	MarketCap(collectionID string, marketplaceID string, from time.Time, to time.Time) ([]datapoint.MarketCap, error)
	Sales(collectionID string, marketplaceID string, from time.Time, to time.Time) ([]datapoint.Sale, error)
	Floor(collectionID string, from time.Time, to time.Time) ([]datapoint.Floor, error)
	Average(collectionID string, from time.Time, to time.Time) ([]datapoint.Average, error)
	Count(collectionID string, from time.Time, to time.Time) ([]datapoint.Count, error)
	MarketplaceUsers(marketplaceID string, from time.Time, to time.Time) ([]datapoint.Users, error)
	NFTPrice(nftID string, from time.Time, to time.Time) ([]datapoint.Price, error)
	NFTAveragePrice(nftID string) (datapoint.Average, error)
}
