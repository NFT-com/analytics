package api

// Value represents the generic datatype for some currency-related stat.
type Value struct {
	ID    string `json:"id,omitempty"`
	Value []Coin `json:"values,omitempty"`
}

// Coin represents a single currency-related stat.
type Coin struct {
	Amount     float64 `json:"amount,omitempty"`
	CurrencyID string  `json:"currency_id,omitempty"`
}

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []Value `json:"data,omitempty"`
}
