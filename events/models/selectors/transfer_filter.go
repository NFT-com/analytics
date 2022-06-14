package selectors

// TransferFilter contains all of the supported query parameters for filtering
// transfer events.
type TransferFilter struct {
	TimestampRange
	HeightRange
	TokenIdentifier

	TransactionHash string `query:"transaction_hash"`
	SenderAddress   string `query:"sender_address"`
	ReceiverAddress string `query:"receiver_address"`
}
