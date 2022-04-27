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

	// Fields related to caching trait and rarity values. `cachemu` is used to lock the struct
	// for access since the GraphQL resolvers are invoked from different goroutines.
	// `cached` is used as a simple check if the values were prefetched.
	// Rarity is derived from trait rarity so it could also be recalculated on the fly
	// and perhaps not even cached.
	cachemu      sync.Mutex `gorm:"-" json:"-"`
	cached       bool       `gorm:"-" json:"-"`
	cachedRarity float64    `gorm:"-" json:"-"`
	cachedTraits []*Trait   `gorm:"-" json:"-"`
}

// CacheTraits caches the trait information so it can be retrieved later
// without doing expensive rarity recomputation.
func (n *NFT) CacheTraits(traits []*Trait) {
	n.cachemu.Lock()
	defer n.cachemu.Unlock()

	n.cachedTraits = traits
	n.cachedRarity = calcRarity(traits)
	n.cached = true
}

// GetCachedRarity retrieves the NFT rarity information from cache, as well as a boolean
// indicating if it was set or not.
func (n *NFT) GetCachedRarity() (float64, bool) {
	n.cachemu.Lock()
	defer n.cachemu.Unlock()

	return n.cachedRarity, n.cached
}

// GetCachedTraits retrieves the traits rarity information from cache, as well as a boolean
// indicating if it was set or not.
func (n *NFT) GetCachedTraits() ([]*Trait, bool) {
	n.cachemu.Lock()
	defer n.cachemu.Unlock()

	return n.cachedTraits, n.cached
}

func calcRarity(traits []*Trait) float64 {

	// Calculate rarity of an NFT by multiplying the ratios of individual traits.
	// For example, if NFT has two traits that are present in 50% of
	// NFTs in a collection, the rarity is 0.5*0.5 = 0.25.
	rarity := 1.0
	for _, trait := range traits {
		rarity = rarity * trait.Rarity
	}

	return rarity
}

// Trait represents a single NFT trait.
// NOTE: `Value` can be an empty string if it represents a trait that the NFT does not have
// (for example when displaying distribution ratio of a rare trait).
type Trait struct {
	Type   string  `json:"type" gorm:"column:name"`
	Value  string  `json:"value" gorm:"column:value"`
	Rarity float64 `json:"rarity" gorm:"column:ratio"`
}
