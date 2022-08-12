package datapoint

import (
	"time"
)

// FIXME: Think of the distinction, where API should switch from addresses to symbols.

// CurrencySnapshot represents the Currency value at a specific date.
type CurrencySnapshot struct {
	Currencies []Currency `json:"currencies,omitempty"`
	Date       time.Time  `json:"date,omitempty"`
}

// Currency has the chain ID and address pair, identifying the fungible token used as payment,
// as well as the amount.
type Currency struct {
	ChainID uint64  `gorm:"column:chain_id" json:"chain_id"`
	Address string  `gorm:"column:currency_address" json:"address"`
	Amount  float64 `gorm:"column:currency_value" json:"amount"`
}
