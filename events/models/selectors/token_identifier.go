package selectors

// TokenIdentifier allows selecting a specific token.
type TokenIdentifier struct {
	ChainID           uint64 `query:"chain_id"`
	CollectionAddress string `query:"collection_address"`
	TokenID           string `query:"token_id"`
}
