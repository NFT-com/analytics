package events

// Mint represents a mint event.
type Mint struct {
	ID          string `gorm:"column:id" json:"-"`
	Collection  string `gorm:"column:collection" json:"collection"`
	BlockNumber string `gorm:"column:block" json:"block"`
	EventIndex  uint   `gorm:"column:event_index" json:"event_index"`
	Transaction string `gorm:"column:transaction_hash" json:"transaction"`
	TokenID     string `gorm:"column:token_id" json:"token_id"`
	Owner       string `gorm:"column:owner" json:"owner"`
	Timestamp   string `gorm:"column:emitted_at" json:"timestamp"`
}

// Transfer represents a transfer event.
type Transfer struct {
	ID          string `gorm:"column:id" json:"-"`
	TokenID     string `gorm:"column:token_id" json:"token_id"`
	Collection  string `gorm:"column:collection" json:"collection"`
	BlockNumber string `gorm:"column:block" json:"block"`
	EventIndex  uint   `gorm:"column:event_index" json:"event_index"`
	Transaction string `gorm:"column:transaction_hash" json:"transaction"`
	From        string `gorm:"column:from_address" json:"from"`
	To          string `gorm:"column:to_address" json:"to"`
	Timestamp   string `gorm:"column:emitted_at" json:"timestamp"`
}

// Sale represents a sale event.
type Sale struct {
	ID          string `gorm:"column:id" json:"-"`
	Marketplace string `gorm:"column:marketplace" json:"marketplace"`
	BlockNumber string `gorm:"column:block" json:"block"`
	EventIndex  uint   `gorm:"column:event_index" json:"event_index"`
	Transaction string `gorm:"column:transaction_hash" json:"transaction"`
	Seller      string `gorm:"column:seller" json:"seller"`
	Buyer       string `gorm:"column:buyer" json:"buyer"`
	Price       string `gorm:"column:price" json:"price"`
	Timestamp   string `gorm:"column:emitted_at" json:"timestamp"`
}

// Burn represents a burn event.
type Burn struct {
	ID          string `gorm:"column:id" json:"-"`
	Collection  string `gorm:"column:collection" json:"collection"`
	BlockNumber string `gorm:"column:block" json:"block"`
	EventIndex  uint   `gorm:"column:event_index" json:"event_index"`
	Transaction string `gorm:"column:transaction_hash" json:"transaction"`
	TokenID     string `gorm:"column:token_id" json:"token_id"`
	Timestamp   string `gorm:"column:emitted_at" json:"timestamp"`
}
