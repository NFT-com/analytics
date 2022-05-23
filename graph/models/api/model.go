package api

// Network represents the blockchain networks.
// Mainnet and testnets of a specific blockchain are distinct network objects.
type Network struct {
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	ID          string `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Address     string `gorm:"column:contract_address" json:"address"`
	Website     string `gorm:"column:website" json:"website"`
	ImageURL    string `gorm:"column:image_url" json:"image_url"`
	NFTs        []*NFT `gorm:"-" json:"nfts"`
	NetworkID   string `gorm:"column:network_id" json:"-"`
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
	TokenID     string   `gorm:"column:token_id" json:"token_id"`
	Owner       string   `gorm:"column:owner" json:"owner"`
	Collection  string   `gorm:"column:collection_id" json:"-"`
	Traits      []*Trait `gorm:"-" json:"traits,omitempty"`
	Rarity      float64  `gorm:"-" json:"rarity,omitempty"`
}

// Trait represents a single NFT trait.
// NOTE: `Value` can be an empty string if it represents a trait that the NFT does not have
// (for example when displaying distribution ratio of a rare trait).
type Trait struct {
	Name   string  `gorm:"column:name" json:"name"`
	Value  string  `gorm:"column:value" json:"value"`
	NFT    string  `gorm:"column:nft_id" json:"-"`
	Rarity float64 `gorm:"-" json:"rarity"`

	// TODO: Add trait type.
}
