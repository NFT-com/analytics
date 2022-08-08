package stats

import (
	"time"
)

// priceResult represents the result of an NFT price query.
type priceResult struct {
	Amount  float64    `gorm:"column:currency_value"`
	Address string     `gorm:"column:currency_address"`
	Time    *time.Time `gorm:"column:emitted_at"`
}

// batchPriceResult represents the result of the batch NFT price query.
type batchPriceResult struct {
	ChainID           uint    `gorm:"column:chain_id"`
	CollectionAddress string  `gorm:"column:collection_address"`
	TokenID           string  `gorm:"column:token_id"`
	TradePrice        float64 `gorm:"column:trade_price"`
}

// batchAveragePriceResult represents the result of the batch NFT average price query.
type batchAveragePriceResult struct {
	ChainID           uint    `gorm:"column:chain_id"`
	CollectionAddress string  `gorm:"column:collection_address"`
	TokenID           string  `gorm:"column:token_id"`
	AveragePrice      float64 `gorm:"column:average_price"`
}

// batchStatResult represents the result of the batch collection volume
// and market cap queries.
type batchStatResult struct {
	ChainID           uint    `gorm:"column:chain_id"`
	CollectionAddress string  `gorm:"column:collection_address"`
	Total             float64 `gorm:"column:total"`
}
