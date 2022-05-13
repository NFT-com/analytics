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
	NFTs        []*NFT `gorm:"-" json:"nfts"`
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
	ID          string   `gorm:"column:id" json:"id"`
	Name        string   `gorm:"column:name" json:"name,omitempty"`
	ImageURL    string   `gorm:"column:image" json:"image_url,omitempty"`
	URI         string   `gorm:"column:uri" json:"uri,omitempty"`
	Description string   `gorm:"column:description" json:"description,omitempty"`
	TokenID     string   `gorm:"column:token_id" json:"tokenID"`
	Owner       string   `gorm:"column:owner" json:"owner"`
	Collection  string   `gorm:"column:collection" json:"-"`
	Traits      []*Trait `gorm:"-" json:"traits,omitempty"`
	Rarity      float64  `gorm:"-" json:"rarity,omitempty"`
}

// Trait represents a single NFT trait.
// NOTE: `Value` can be an empty string if it represents a trait that the NFT does not have
// (for example when displaying distribution ratio of a rare trait).
type Trait struct {
	Type   string  `gorm:"column:name" json:"type"`
	Value  string  `gorm:"column:value" json:"value"`
	Rarity float64 `gorm:"column:ratio" json:"rarity"`
	NFT    string  `gorm:"column:nft" json:"-"`
}
