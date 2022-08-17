package events

import (
	"time"
)

type Sale struct {
	ID                 string    `gorm:"column:id" json:"id"`
	ChainID            uint64    `gorm:"column:chain_id" json:"chain_id"`
	MarketplaceAddress string    `gorm:"column:marketplace_address" json:"marketplace_address"`
	CollectionAddress  string    `gorm:"column:collection_address" json:"collection_address"`
	TokenID            string    `gorm:"column:token_id" json:"token_id"`
	BlockNumber        uint64    `gorm:"column:block_number" json:"block_number"`
	TransactionHash    string    `gorm:"column:transaction_hash" json:"transaction_hash"`
	EventIndex         uint      `gorm:"column:event_index" json:"event_index"`
	SellerAddress      string    `gorm:"column:seller_address" json:"seller_address"`
	BuyerAddress       string    `gorm:"column:buyer_address" json:"buyer_address"`
	TokenCount         float64   `gorm:"column:token_count" json:"token_count"`
	CurrencyAddress    string    `gorm:"column:currency_address" json:"currency_address"`
	CurrencyValue      float64   `gorm:"column:currency_value" json:"currency_value"`
	EmittedAt          time.Time `gorm:"column:emitted_at" json:"emitted_at"`
}
