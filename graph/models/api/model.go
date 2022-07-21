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
	// NFTs        []*NFT  `gorm:"-" json:"nfts"`
	NetworkID string        `gorm:"column:network_id" json:"-"`
	Volume    float64       `gorm:"-" json:"volume"`
	MarketCap float64       `gorm:"-" json:"market_cap"`
	Sales     uint64        `gorm:"-" json:"sales"`
	NFTs      NFTConnection `gorm:"-" json:"nfts"`
}

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string  `gorm:"column:id" json:"id"`
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Website     string  `gorm:"column:website" json:"website"`
	Volume      float64 `gorm:"-" json:"volume"`
	MarketCap   float64 `gorm:"-" json:"market_cap"`
	Sales       uint64  `gorm:"-" json:"sales"`
	Users       uint64  `gorm:"-" json:"users"`
}

// NFT represents a single Non-Fungible Token.
type NFT struct {
	ID           string  `gorm:"column:id" json:"id"`
	Name         string  `gorm:"column:name" json:"name,omitempty"`
	ImageURL     string  `gorm:"column:image" json:"image_url,omitempty"`
	URI          string  `gorm:"column:uri" json:"uri,omitempty"`
	Description  string  `gorm:"column:description" json:"description,omitempty"`
	TokenID      string  `gorm:"column:token_id" json:"token_id"`
	Collection   string  `gorm:"column:collection_id" json:"-"`
	Owners       []Owner `gorm:"-" json:"owners,omitempty"`
	Traits       []Trait `gorm:"-" json:"traits,omitempty"`
	Rarity       float64 `gorm:"-" json:"rarity,omitempty"`
	TradingPrice float64 `gorm:"-" json:"trading_price,omitempty"`
	AveragePrice float64 `gorm:"-" json:"average_price,omitempty"`
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

// FIXME: Add documentation.

type NFTConnection struct {
	Edges    []NFTEdge `json:"edges"`
	PageInfo PageInfo  `json:"pageInfo"`
}

type NFTEdge struct {
	Node   *NFT   `json:"node"`
	Cursor string `json:"cursor"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
}
