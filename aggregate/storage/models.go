package storage

type collectionAddress struct {
	ChainID uint   `gorm:"column:chain_id"`
	Address string `gorm:"column:contract_address"`
}
