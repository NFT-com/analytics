package events

// Transfer represents a transfer event.
type Transfer struct {
	ID                string `gorm:"column:id" json:"-"`
	ChainID           string `gorm:"column:chain_id" json:"chain_id"`
	CollectionAddress string `gorm:"column:collection_address" json:"collection_addresss"`
	TokenID           string `gorm:"column:token_id" json:"token_id"`
	BlockNumber       string `gorm:"column:block_number" json:"block_number"`
	EventIndex        uint   `gorm:"column:event_index" json:"event_index"`
	Transaction       string `gorm:"column:transaction_hash" json:"transaction"`
	Sender            string `gorm:"column:sender_address" json:"sender_address"`
	Receiver          string `gorm:"column:receiver_address" json:"receiver_address"`
	Timestamp         string `gorm:"column:emitted_at" json:"timestamp"`
}

// Sale represents a sale event.
type Sale struct {
	ID                 string `gorm:"column:id" json:"-"`
	ChainID            string `gorm:"column:chain_id" json:"chain_id"`
	CollectionAddress  string `gorm:"column:collection_address" json:"collection_addresss"`
	TokenID            string `gorm:"column:token_id" json:"token_id"`
	MarketplaceAddress string `gorm:"column:marketplace_address" json:"marketplace_address"`
	BlockNumber        string `gorm:"column:block_number" json:"block_number"`
	EventIndex         uint   `gorm:"column:event_index" json:"event_index"`
	Transaction        string `gorm:"column:transaction_hash" json:"transaction"`
	Seller             string `gorm:"column:seller_address" json:"seller_address"`
	Buyer              string `gorm:"column:buyer_address" json:"buyer_address"`
	Price              string `gorm:"column:trade_price" json:"trade_price"`
	Timestamp          string `gorm:"column:emitted_at" json:"timestamp"`
}
