package lookup

type networkAddress struct {
	ID              string `gorm:"column:id"`
	ChainID         uint64 `gorm:"column:chain_id"`
	ContractAddress string `gorm:"column:contract_address"`
}

type nftIdentifier struct {
	ID      string `gorm:"column:id"`
	ChainID uint64 `gorm:"column:chain_id"`
	Address string `gorm:"column:contract_address"`
	TokenID string `gorm:"column:token_id"`
}
