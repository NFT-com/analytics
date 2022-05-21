package lookup

type networkAddress struct {
	ChainID         uint   `gorm:"column:chain_id"`
	ContractAddress string `gorm:"column:contract_address"`
}

type nftIdentifier struct {
	ChainID uint   `gorm:"column:chain_id"`
	Address string `gorm:"column:contract_address"`
	TokenID string `gorm:"column:token_id"`
}