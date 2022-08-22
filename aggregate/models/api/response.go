package api

import (
	"time"
)

// Value represents the generic datatype for some stat measured in fungible tokens,
// such as NFT price, trading volume of a collection etc.
type Value struct {
	ID    string `json:"id"`
	Value []Coin `json:"value"`
}

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []Value `json:"data,omitempty"`
}

// Coin represents an amount of fungible tokens, consisting of the number of tokens
// and the currency identifier.
type Coin struct {
	CurrencyID string  `json:"currency_id"`
	Value      float64 `json:"value"`
}

// CoinSnapshot represents the value of a stat measured in fungible tokens at
// a specific point in time, for example - the price of an NFT in a sale.
type CoinSnapshot struct {
	Value []Coin     `json:"value"`
	Time  *time.Time `json:"timestamp,omitempty"`
}

// ValueHistory represents historic values of a stat measured in fungible tokens.
type ValueHistory struct {
	ID        string         `json:"id"`
	Snapshots []CoinSnapshot `json:"snapshots"`
}
