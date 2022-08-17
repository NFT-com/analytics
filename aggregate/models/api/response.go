package api

import (
	"time"
)

// Value represents the generic datatype for some currency-related stat.
type Value struct {
	ID    string `json:"id,omitempty"`
	Value []Coin `json:"values,omitempty"`
}

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []Value `json:"data,omitempty"`
}

// Coin represents a single currency-related stat.
type Coin struct {
	CurrencyID string  `json:"currency_id,omitempty"`
	Value      float64 `json:"value,omitempty"`
}

// CoinSnapshot represents a currency-related stat at a certain point in time.
type CoinSnapshot struct {
	CurrencyID string     `json:"currency_id,omitempty"`
	Time       *time.Time `json:"timestamp,omitempty"`
	Value      float64    `json:"value,omitempty"`
}

// PriceHistory represents the NFT prices at different points in time.
type PriceHistory struct {
	ID     string         `json:"id,omitempty"`
	Prices []CoinSnapshot `json:"prices"`
}
