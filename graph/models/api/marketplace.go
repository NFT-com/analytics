package api

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	ID          string     `gorm:"column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Website     string     `gorm:"column:website" json:"website"`
	Volume      []Currency `gorm:"-" json:"volume"`
	MarketCap   []Currency `gorm:"-" json:"market_cap"`
	Sales       uint64     `gorm:"-" json:"sales"`
	Users       uint64     `gorm:"-" json:"users"`
}
