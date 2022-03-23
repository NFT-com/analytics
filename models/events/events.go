package events

// FIXME: Check - plural or singular for package name?
// FIXME: Types modeled after tentative_tables.sql and are subject to change.
// FIXME: Block numbers should be larger than int64

// Mint represents a mint event.
type Mint struct {
	TokenID     string `json:"token_id" gorm:"column:id"`
	Chain       string `json:"chain" gorm:"column:chain_id"`
	Collection  string `json:"collection" gorm:"column:collection_id"`
	BlockNumber int64  `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	Minter      string `json:"minter" gorm:"column:minter"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}

// Transfer represents a transfer event.
type Transfer struct {
	TokenID     string `json:"token_id" gorm:"column:id"`
	Chain       string `json:"chain" gorm:"column:chain_id"`
	Collection  string `json:"collection" gorm:"column:collection_id"`
	BlockNumber int64  `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	From        string `json:"from" gorm:"column:from_address"`
	To          string `json:"to" gorm:"column:to_address"`
	Price       string `json:"price" gorm:"column:price"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}

// Sale represents a sale event.
type Sale struct {
	TokenID     string `json:"token_id" gorm:"column:id"`
	Chain       string `json:"chain" gorm:"column:chain_id"`
	Collection  string `json:"collection" gorm:"column:collection_id"`
	BlockNumber int64  `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	Marketplace string `json:"marketplace" gorm:"column:marketplace"`
	Owner       string `json:"owner" gorm:"column:owner"`
	Price       string `json:"price" gorm:"column:price"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}

// Burn represents a burn event.
type Burn struct {
	TokenID     string `json:"token_id" gorm:"column:id"`
	Chain       string `json:"chain" gorm:"column:chain_id"`
	Collection  string `json:"collection" gorm:"column:collection_id"`
	BlockNumber int64  `json:"block" gorm:"column:block"`
	Transaction string `json:"transaction" gorm:"column:transaction_hash"`
	Burner      string `json:"burner" gorm:"column:burner"`
	Timestamp   string `json:"timestamp" gorm:"column:emitted_at"`
}
