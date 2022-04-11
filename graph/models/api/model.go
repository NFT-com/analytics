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
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name,omitempty"`
	ImageURL    string `gorm:"column:image" json:"image_url,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	TokenID     string `gorm:"column:token_id" json:"tokenID"`
	Owner       string `gorm:"column:owner" json:"owner"`
	Collection  string `gorm:"column:collection" json:"-"`
}

// Trait represents a single NFT trait.
type Trait struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

// Trait ratio represents the ratio of NFTs in a collection with this specific trait.
type TraitRatio struct {
	Trait Trait   `json:"trait"`
	Ratio float64 `json:"ratio"`
}
