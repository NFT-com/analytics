package datapoint

import (
	"time"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// CoinSnapshot represents the value of a metric at a specific date, expressed in coins.
type CoinSnapshot struct {
	Coins []Coin    `json:"coins,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}

// Coin represents a pairing of a currency and an amount.
type Coin struct {
	Currency identifier.Currency `json:"currency"`
	Amount   float64             `json:"amount"`
}
