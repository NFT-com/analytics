package api

// Chain represents the chain and its networks.
type Chain struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Marketplaces []*Marketplace `json:"marketplaces"`
	Collections  []*Collection  `json:"collections"`
}

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Address      string         `json:"address"`
	Chain        *Chain         `json:"chain"`
	Marketplaces []*Marketplace `json:"marketplaces"`
	Nfts         []*NFT         `json:"nfts"`
}

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Chains      []*Chain      `json:"chains"`
	Collections []*Collection `json:"collections"`
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

// TableName returns the name of the underlying database table for the NFT.
func (n *NFT) TableName() string {
	return "nft"
}

// TableName returns the name of the underlying database table for the Collection.
func (c *Collection) TableName() string {
	return "collection"
}
