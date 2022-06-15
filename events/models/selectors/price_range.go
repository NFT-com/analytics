package selectors

import (
	"math/big"
)

type PriceRange struct {
	StartPrice big.Int `query:"start_price"`
	EndPrice   big.Int `query:"end_price"`
}
