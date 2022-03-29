package events

// MintSelector contains all of the supported query parameters for filtering
// mint events.
type MintSelector struct {
	TimeSelector
	TokenSelector

	Transaction string `query:"transaction"`
	Owner       string `query:"owner"`
}

// TransferSelector contains all of the supported query parameters for filtering
// transfer events.
type TransferSelector struct {
	TimeSelector
	TokenSelector

	Transaction string `query:"transaction"`
	From        string `query:"from"`
	To          string `query:"to"`
}

// SaleSelector contains all of the supported query parameters for filtering
// sale events.
type SaleSelector struct {
	TimeSelector

	Transaction string `query:"transaction"`
	Marketplace string `query:"marketplace"`
	Seller      string `query:"seller"`
	Buyer       string `query:"buyer"`
	Price       string `query:"price"`
}

// BurnSelector contains all of the supported query paramters for filtering
// burn events.
type BurnSelector struct {
	TimeSelector
	TokenSelector

	Transaction string `query:"transaction"`
}

// TokenSelector allows selecting a specific token.
type TokenSelector struct {
	Collection string `query:"collection"`
	TokenID    string `query:"token_id"`
}

// TimeSelector allows selecting events in a time range.
type TimeSelector struct {
	Start string `query:"start"`
	End   string `query:"end"`
}

// FIXME: Validate format if parameters are set.
// FIXME: Add filters for start and end height.
// FIXME: Change start/end timestamps.
