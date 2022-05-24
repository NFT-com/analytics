package stats

// batchPriceResult represents the result of the batch NFT price query.
type batchPriceResult struct {
	ChainID           uint   `gorm:"column:chain_id"`
	CollectionAddress string `gorm:"column:collection_address"`
	TokenID           string `gorm:"column:token_id"`
	/// FIXME: Change type.
	TradePrice string `gorm:"column:trade_price"`
}
