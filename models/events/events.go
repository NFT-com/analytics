package events

// FIXME: Check - plural or singular for package name?
// FIXME: Types modeled after tentative_tables.sql and are subject to change.
// FIXME: Block numbers should be larger than int64

// Mint represents a mint event.
type Mint struct {
	TokenID     string `json:"token_id" gorm:"id"`
	Chain       string `json:"chain" gorm:"chain_id"`
	Collection  string `json:"collection" gorm:"collection_id"`
	BlockNumber int64  `json:"block" gorm:"block"`
	Transaction string `json:"transaction" gorm:"transaction_hash"`
	Minter      string `json:"minter" gorm:"minter"`
	Timestamp   string `json:"timestamp" gorm:"emitted_at"`
}

// Transfer represents a transfer event.
type Transfer struct {
	TokenID     string `json:"token_id" gorm:"id"`
	Chain       string `json:"chain" gorm:"chain_id"`
	Collection  string `json:"collection" gorm:"collection_id"`
	BlockNumber int64  `json:"block" gorm:"block"`
	Transaction string `json:"transaction" gorm:"transaction_hash"`
	From        string `json:"from" gorm:"from_address"`
	To          string `json:"to" gorm:"to_address"`
	Price       string `json:"price" gorm:"price"`
	Timestamp   string `json:"timestamp" gorm:"emitted_at"`
}

// Sale represents a sale event.
type Sale struct {
	TokenID     string `json:"token_id" gorm:"id"`
	Chain       string `json:"chain" gorm:"chain_id"`
	Collection  string `json:"collection" gorm:"collection_id"`
	BlockNumber int64  `json:"block" gorm:"block"`
	Transaction string `json:"transaction" gorm:"transaction_hash"`
	Marketplace string `json:"marketplace" gorm:"marketplace"`
	Owner       string `json:"owner" gorm:"owner"`
	Price       string `json:"price" gorm:"price"`
	Timestamp   string `json:"timestamp" gorm:"emitted_at"`
}

// Burn represents a burn event.
type Burn struct {
	TokenID     string `json:"token_id" gorm:"id"`
	Chain       string `json:"chain" gorm:"chain_id"`
	Collection  string `json:"collection" gorm:"collection_id"`
	BlockNumber int64  `json:"block" gorm:"block"`
	Transaction string `json:"transaction" gorm:"transaction_hash"`
	Burner      string `json:"burner" gorm:"burner"`
	Timestamp   string `json:"timestamp" gorm:"emitted_at"`
}
