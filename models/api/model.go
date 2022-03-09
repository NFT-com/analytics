package api

// Chain represents the chain and its networks.
type Chain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// FIXME: Add marketplaces link
	// FIXME: Add collections link
}

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	ChainID     string `json:"-"`
	// FIXME: Add nfts link
	// FIXME: Add marketplace link
	// FIXME: Add chain link
}

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// FIXME: Add chains link
	// FIXME: Add collections link
}

type NFT struct {
	ID           string  `json:"id"`
	TokenID      string  `json:"tokenID"`
	Owner        string  `json:"owner"`
	URI          string  `json:"uri"`
	Rarity       float64 `json:"rarity"`
	CollectionID string  `json:"-"`
}

// FIXME: think if you need/want another layer to this - storage specific functions below.

// TableName returns the name of the underlying database table for the Chain.
func (c *Chain) TableName() string {
	return "chain"
}

// TableName returns the name of the underlying database table for the Collection.
func (c *Collection) TableName() string {
	return "collection"
}

// TableName returns the name of the underlying database table for the NFT.
func (n *NFT) TableName() string {
	return "nft"
}
