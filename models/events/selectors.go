package events

// MintSelector contains all of the supported query parameters for filtering
// mint events.
type MintSelector struct {
	Collection  string `query:"collection"`
	TokenID     string `query:"token_id"`
	Transaction string `query:"transaction"`
	Owner       string `query:"owner"`
	Start       string `query:"start"`
	End         string `query:"end"`
}

// TransferSelector contains all of the supported query parameters for filtering
// transfer events.
type TransferSelector struct {
	Collection  string `query:"collection"`
	TokenID     string `query:"token_id"`
	Transaction string `query:"transaction"`
	From        string `query:"from"`
	To          string `query:"to"`
	Start       string `query:"start"`
	End         string `query:"end"`
}

// SaleSelector contains all of the supported query parameters for filtering
// sale events.
type SaleSelector struct {
	Transaction string `query:"transaction"`
	Marketplace string `query:"marketplace"`
	Seller      string `query:"seller"`
	Buyer       string `query:"buyer"`
	Price       string `query:"price"`
	Start       string `query:"start"`
	End         string `query:"end"`
}

// BurnSelector contains all of the supported query paramters for filtering
// burn events.
type BurnSelector struct {
	Collection  string `query:"collection"`
	TokenID     string `query:"token_id"`
	Transaction string `query:"transaction"`
	Start       string `query:"start"`
	End         string `query:"end"`
}

// FIXME: Implement remaining selectors.
// FIXME: Pick a good place for event selectors.
// FIXME: Validate format if parameters are set.
// FIXME: Think of event-specific filters.
// FIXME: Add filters for start and end height.
// FIXME: Change start/end timestamps.
