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

// Currency has the amount and the address of the fungible token used as payment
// in a transaction.
type Currency struct {
	Amount  float64 `gorm:"column:currency_value" json:"amount"`
	Address string  `gorm:"column:currency_address" json:"address"`
}
