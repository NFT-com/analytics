package identifier

// Address identifier represents a single address on a blockchain.
type Address struct {
	ChainID uint
	Address string
}

// NFT identifier represents a single NFT on a blockchain.
type NFT struct {
	Collection Address
	TokenID    string
}
