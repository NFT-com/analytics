package api

// Chain represents the chain and its networks.
type Chain struct {
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Address     string `gorm:"column:address" json:"address"`
	Website     string `gorm:"column:website" json:"website"`
	ImageURL    string `gorm:"column:image_url" json:"image_url"`
	TokenURI    string `gorm:"column:uri" json:"token_uri"`
	ChainID     string `gorm:"column:chain_id" json:"-"`
}

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Website     string `gorm:"column:website" json:"website"`
}

// NFT represents a single Non-Fungible Token.
type NFT struct {
	ID         string  `gorm:"column:id" json:"id"`
	TokenID    string  `gorm:"column:token_id" json:"tokenID"`
	Owner      string  `gorm:"column:owner" json:"owner"`
	Rarity     float64 `gorm:"column:rarity" json:"rarity"`
	Collection string  `gorm:"column:collection" json:"-"`
}