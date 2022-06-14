package selectors

// SalesFilter contains all of the supported query parameters for filtering
// sale events.
type SalesFilter struct {
	TimestampRange
	HeightRange
	PriceRange

	ChainID            uint64 `query:"chain_id"`
	MarketplaceAddress string `query:"marketplace_address"`
	CollectionAddress  string `query:"collection_address"`
	TokenID            string `query:"token_id"`
	TransactionHash    string `query:"transaction_address"`
	SellerAddress      string `query:"seller_address"`
	BuyerAddress       string `query:"buyer_address"`
}
