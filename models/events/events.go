package events

// FIXME: Check - plural or singular for package name?
// FIXME: Types modeled after tentative_tables.sql and are subject to change.
// FIXME: Use numbers for block numbers.

// Mint represents a mint event.
type Mint struct {
	ID          string `gorm:"column:id" json:"-"`
	Collection  string `gorm:"column:collection" json:"collection"`
	BlockNumber string `gorm:"column:block" json:"block"`
	Transaction string `gorm:"column:transaction_hash" json:"transaction"`
	TokenID     string `gorm:"column:nft_id" json:"token_id"`
	Owner       string `gorm:"column:owner" json:"owner"`
	Timestamp   string `gorm:"column:emitted_at" json:"timestamp"`
}

// Transfer represents a transfer event.
type Transfer struct {
	ID          string `gorm:"column:id" json:"-"`
	TokenID     string `json:"token_id" gorm:"column:nft_id"`
	Collection  string `json:"collection" gorm:"column:collection"`
	BlockNumber string `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	From        string `json:"from" gorm:"column:from_address"`
	To          string `json:"to" gorm:"column:to_address"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}

// Sale represents a sale event.
type Sale struct {
	ID          string `gorm:"column:id" json:"-"`
	Marketplace string `json:"marketplace" gorm:"column:marketplace"`
	BlockNumber string `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	Seller      string `json:"seller" gorm:"column:seller"`
	Buyer       string `json:"buyer" gorm:"column:buyer"`
	Price       string `json:"price" gorm:"column:price"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}

// Burn represents a burn event.
type Burn struct {
	ID          string `gorm:"column:id" json:"-"`
	Collection  string `json:"collection" gorm:"column:collection"`
	BlockNumber string `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	TokenID     string `json:"token_id" gorm:"column:nft_id"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}
