package api

// BatchResponse represents the API response for batch stats request.
type BatchResponse struct {
	Data []StatValue `json:"data,omitempty"`
}

// StatValue has the requested stat for an entity (e.g. a collection or marketplace) in a batch request.
type StatValue struct {
	ID    string  `json:"id,omitempty"`
	Value float64 `json:"value,omitempty"`
}
