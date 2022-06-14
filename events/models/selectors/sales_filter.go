package selectors

// SalesFilter contains all of the supported query parameters for filtering
// sale events.
type SalesFilter struct {
	TimestampRange
	HeightRange
	TokenIdentifier

	MarketplaceAddress string `query:"marketplace_address"`
	CollectionAddress  string `query:"collection_address"`
	TransactionHash    string `query:"transaction_address"`
	SellerAddress      string `query:"seller_address"`
	BuyerAddress       string `query:"buyer_address"`
	TradePrice         string `query:"trade_price"`
}
