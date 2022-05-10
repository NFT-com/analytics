package events

// TransferSelector contains all of the supported query parameters for filtering
// transfer events.
type TransferSelector struct {
	TimeSelector
	BlockSelector
	TokenSelector

	Transaction string `query:"transaction"`
	Sender      string `query:"sender"`
	Receiver    string `query:"receiver"`
}

// SaleSelector contains all of the supported query parameters for filtering
// sale events.
type SaleSelector struct {
	TimeSelector
	BlockSelector
	TokenSelector

	Marketplace string `query:"marketplace"`
	Transaction string `query:"transaction"`
	Seller      string `query:"seller"`
	Buyer       string `query:"buyer"`
	Price       string `query:"price"`
}

// TokenSelector allows selecting a specific token.
type TokenSelector struct {
	Chain      string `query:"chain"`
	Collection string `query:"collection"`
	TokenID    string `query:"token_id"`
}
