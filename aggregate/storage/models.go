package storage

type collectionAddress struct {
	ChainID uint   `gorm:"column:chain_id"`
	Address string `gorm:"column:contract_address"`
}

type nftIdentifier struct {
	collectionAddress
	TokenID string `gorm:"column:token_id"`
}
