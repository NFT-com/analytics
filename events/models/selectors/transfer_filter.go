package selectors

// TransferFilter contains all of the supported query parameters for filtering
// transfer events.
type TransferFilter struct {
	TimestampRange
	HeightRange

	ChainID           uint64 `query:"chain_id"`
	CollectionAddress string `query:"collection_address"`
	TokenID           string `query:"token_id"`
	TransactionHash   string `query:"transaction_hash"`
	SenderAddress     string `query:"sender_address"`
	ReceiverAddress   string `query:"receiver_address"`
}
