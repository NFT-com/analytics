package stats

import (
	"time"
)

// priceResult represents the result of an NFT price query.
type priceResult struct {
	ChainID uint64     `gorm:"column:chain_id"`
	Address string     `gorm:"column:currency_address"`
	Amount  float64    `gorm:"column:currency_value"`
	Time    *time.Time `gorm:"column:emitted_at"`
}

// batchPriceResult represents the result of the batch NFT price query.
type batchPriceResult struct {
	ChainID           uint64  `gorm:"column:chain_id"`
	CollectionAddress string  `gorm:"column:collection_address"`
	TokenID           string  `gorm:"column:token_id"`
	CurrencyAmount    float64 `gorm:"column:currency_value"`
	CurrencyAddress   string  `gorm:"column:currency_address"`
}

// batchStatResult represents the result of the batch collection volume
// and market cap queries.
type batchStatResult struct {
	ChainID           uint64  `gorm:"column:chain_id"`
	CollectionAddress string  `gorm:"column:collection_address"`
	Amount            float64 `gorm:"column:currency_value"`
	Address           string  `gorm:"column:currency_address"`
}

// lowestPriceResult represents the result of the SQL query for the lowest price.
type lowestPriceResult struct {
	ChainID uint64  `gorm:"column:chain_id"`
	Address string  `gorm:"column:currency_address"`
	Amount  float64 `gorm:"column:currency_value"`
	Start   string  `gorm:"column:start_date"`
	End     string  `gorm:"column:end_date"`
}

// datedPriceResult represents the results of queries returning a stat (e.g. volume, market_cap)
// at a certain date point.
type datedPriceResult struct {
	ChainID uint64    `gorm:"column:chain_id"`
	Address string    `gorm:"column:currency_address"`
	Amount  float64   `gorm:"column:currency_value"`
	Date    time.Time `gorm:"column:date"`
}
