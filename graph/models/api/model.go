package api

import (
	"sync"
)

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
	URI         string `gorm:"column:uri" json:"uri,omitempty"`
	Description string `gorm:"column:description" json:"description,omitempty"`
	TokenID     string `gorm:"column:token_id" json:"tokenID"`
	Owner       string `gorm:"column:owner" json:"owner"`
	Collection  string `gorm:"column:collection" json:"-"`

	// Fields related to caching rarity values. `cachemu` is used to lock the struct
	// for access since the GraphQL resolvers are invoked from different goroutines.
	// `cached` is used as a simple check if the values were prefetched.
	cachemu      sync.Mutex `gorm:"-" json:"-"`
	cached       bool       `gorm:"-" json:"-"`
	cachedRarity float64    `gorm:"-" json:"-"`
}

// CacheRarity caches the rarity for the specific NFT so it can be retrieved later
// without doing expensive recomputation.
func (n *NFT) CacheRarity(rarity float64) {
	n.cachemu.Lock()
	defer n.cachemu.Unlock()

	n.cachedRarity = rarity
	n.cached = true
}

// GetCachedRarity retrieves the NFT rarity information from cache, as well as a boolean
// indicating if it was set or not.
func (n *NFT) GetCachedRarity() (float64, bool) {
	n.cachemu.Lock()
	defer n.cachemu.Unlock()

	return n.cachedRarity, n.cached
}

// Trait represents a single NFT trait.
// NOTE: `Value` can be an empty string if it represents a trait that the NFT does not have
// (for example when displaying distribution ratio of a rare trait).
type Trait struct {
	Type   string  `json:"type" gorm:"column:name"`
	Value  string  `json:"value" gorm:"column:value"`
	Rarity float64 `json:"rarity" gorm:"column:ratio"`
}
