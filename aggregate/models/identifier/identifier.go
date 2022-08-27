package identifier

const (
	ZeroAddress = "0x0000000000000000000000000000000000000000"
	DeadAddress = "0x000000000000000000000000000000000000dEaD"
)

// Address identifier represents a single address on a blockchain.
type Address struct {
	ChainID uint64
	Address string
}

// NFT identifier represents a single NFT on a blockchain.
type NFT struct {
	Collection Address
	TokenID    string
}

// Currency represents the chain ID and address pair, identifying a fungible token used as payment.
type Currency struct {
	ChainID uint64 `gorm:"column:chain_id" json:"chain_id"`
	Address string `gorm:"column:currency_address" json:"address"`
}
