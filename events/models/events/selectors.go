package events

// TransferSelector contains all of the supported query parameters for filtering
// transfer events.
type TransferSelector struct {
	TimeSelector
	BlockSelector
	TokenSelector

	Transaction string `query:"transaction"`
	From        string `query:"from"`
	To          string `query:"to"`
}

// SaleSelector contains all of the supported query parameters for filtering
// sale events.
type SaleSelector struct {
	TimeSelector
	BlockSelector

	Transaction string `query:"transaction"`
	Marketplace string `query:"marketplace"`
	Seller      string `query:"seller"`
	Buyer       string `query:"buyer"`
	Price       string `query:"price"`
}

// TokenSelector allows selecting a specific token.
type TokenSelector struct {
	Collection string `query:"collection"`
	TokenID    string `query:"token_id"`
}
