package datapoint

// FIXME: Think of the distinction, where API should switch from addresses to symbols.

// Currency has the amount and the address of the fungible token used as payment
// in a transaction.
type Currency struct {
	Amount  float64 `gorm:"column:currency_value" json:"amount"`
	Address string  `gorm:"column:currency_address" json:"address"`
}
