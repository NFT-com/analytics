package api

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	ID          string        `gorm:"column:id" json:"id"`
	Name        string        `gorm:"column:name" json:"name"`
	Description string        `gorm:"column:description" json:"description"`
	Address     string        `gorm:"column:contract_address" json:"address"`
	Website     string        `gorm:"column:website" json:"website"`
	ImageURL    string        `gorm:"column:image_url" json:"image_url"`
	NetworkID   string        `gorm:"column:network_id" json:"-"`
	Volume      []Currency    `gorm:"-" json:"volume"`
	MarketCap   []Currency    `gorm:"-" json:"market_cap"`
	Sales       uint64        `gorm:"-" json:"sales"`
	NFTs        NFTConnection `gorm:"-" json:"nfts"`
}

// Currency represents a fungible token, typically used for payment.
// FIXME: Check - this is called 'Coin' in the aggregation API.
type Currency struct {
	Amount float64 `json:"amount"`
	Symbol string  `json:"symbol"`
}

// NFTConnection is used for paginated access to the Collection NFT list.
type NFTConnection struct {
	Edges    []NFTEdge `json:"edges"`
	PageInfo PageInfo  `json:"pageInfo"`
}

// NFTEdge contains the NFT data, as well as pagination-related metadata.
type NFTEdge struct {
	Node   *NFT   `json:"node"`
	Cursor string `json:"cursor"`
}

// PageInfo contains the information related to the pagination end condition,
// as well as the cursor that can be used to restart pagination.
// Note that pagination can be started from the beginning simply by omitting the
// cursor entirely.
type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
}
