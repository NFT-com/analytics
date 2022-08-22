package api

import (
	"time"
)

// Value represents the generic datatype for some currency-related stat.
type Value struct {
	ID    string `json:"id"`
	Value []Coin `json:"value"`
}

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []Value `json:"data,omitempty"`
}

// Coin represents a single currency-related stat.
type Coin struct {
	CurrencyID string  `json:"currency_id"`
	Value      float64 `json:"value"`
}

// CoinSnapshot represents a currency-related stat at a certain point in time.
type CoinSnapshot struct {
	Value []Coin     `json:"value"`
	Time  *time.Time `json:"timestamp,omitempty"`
}

// ValueHistory represents historic values of a currency-related stat.
type ValueHistory struct {
	ID        string         `json:"id"`
	Snapshots []CoinSnapshot `json:"snapshots"`
}
