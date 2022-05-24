package api

// NFTPriceResponse represents the API response for the NFT batch price request.
type NFTPriceResponse struct {
	Prices []NFTPrice `json:"prices,omitempty"`
}

// NFTPrice represents the single NFT price record.
type NFTPrice struct {
	ID string `json:"id,omitempty"`
	// FIXME: Should not be a string.
	Price string `json:"price,omitempty"`
}
