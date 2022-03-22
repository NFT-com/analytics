package events

import (
	"math/big"
)

// FIXME: Check these types.

// Mint represents a mint event.
type Mint struct {
	Chain       string
	Collection  string
	TokenID     big.Int
	Minter      string
	BlockNumber big.Int
	Timestamp   string
}

// Transfer represents a transfer event.
type Transfer struct {
	Chain       string
	Collection  string
	TokenID     big.Int
	From        string
	To          string
	Price       string
	BlockNumber big.Int
	Timestamp   string
}

// Sale represents a sale event.
type Sale struct {
	Chain       string
	Collection  string
	TokenID     big.Int
	BlockNumber big.Int
	Timestamp   string
}

// Burn represents a burn event.
type Burn struct {
	Chain       string
	Collection  string
	TokenID     big.Int
	Owner       string
	BlockNumber big.Int
	Timestamp   string
}
