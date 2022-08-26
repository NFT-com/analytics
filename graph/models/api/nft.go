package api

// NFT represents a single Non-Fungible Token.
type NFT struct {
	ID           string     `gorm:"column:id" json:"id"`
	Name         string     `gorm:"column:name" json:"name,omitempty"`
	ImageURL     string     `gorm:"column:image" json:"image_url,omitempty"`
	URI          string     `gorm:"column:uri" json:"uri,omitempty"`
	Description  string     `gorm:"column:description" json:"description,omitempty"`
	TokenID      string     `gorm:"column:token_id" json:"token_id"`
	Collection   string     `gorm:"column:collection_id" json:"-"`
	Deleted      bool       `gorm:"column:deleted" json:"-"`
	Owners       []Owner    `gorm:"-" json:"owners,omitempty"`
	Traits       []Trait    `gorm:"-" json:"traits,omitempty"`
	Rarity       float64    `gorm:"-" json:"rarity,omitempty"`
	TradingPrice []Currency `gorm:"-" json:"trading_price,omitempty"`
	AveragePrice []Currency `gorm:"-" json:"average_price,omitempty"`
}

// Owner reprecsents the owner of the NFT, along with the information of how many tokens it has.
type Owner struct {
	Address string `gorm:"column:owner" json:"address"`
	NFTID   string `gorm:"column:nft_id" json:"-"`
	Number  int    `gorm:"column:number" json:"number"`
}

// Trait represents a single NFT trait.
// NOTE: `Value` can be an empty string if it represents a trait that the NFT does not have
// (for example when displaying distribution ratio of a rare trait).
type Trait struct {
	Name   string  `gorm:"column:name" json:"name"`
	Value  string  `gorm:"column:value" json:"value,omitempty"`
	Type   string  `gorm:"column:type" json:"type,omitempty"`
	NFT    string  `gorm:"column:nft_id" json:"-"`
	Rarity float64 `gorm:"-" json:"rarity"`
	// TODO: Add trait type handling when implemented.
	// https://github.com/NFT-com/analytics/issues/27
}
