package stats

// batchPriceResult represents the result of the batch NFT price query.
type batchPriceResult struct {
	ChainID           uint   `gorm:"column:chain_id"`
	CollectionAddress string `gorm:"column:collection_address"`
	TokenID           string `gorm:"column:token_id"`
	/// FIXME: Change type.
	TradePrice string `gorm:"column:trade_price"`
}

type batchVolumeResult struct {
	ChainID           uint   `gorm:"column:chain_id"`
	CollectionAddress string `gorm:"column:collection_address"`
	// FIXME: Change type
	Total string `gorm:"column:total"`
}

// FIXME: Could be modeled with the same type as volume.
type batchMarketCapResult struct {
	ChainID           uint   `gorm:"column:chain_id"`
	CollectionAddress string `gorm:"column:collection_address"`
	// FIXME: Change type
	Total string `gorm:"column:total"`
}
