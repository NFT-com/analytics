package stats

import (
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// convertPricesToCoins will take an array of priceResult records and create
// the corresponding datapoint.Coin types.
func convertPricesToCoins(prices []priceResult) []datapoint.Coin {

	coins := make([]datapoint.Coin, 0, len(prices))
	for _, p := range prices {

		volume := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: p.ChainID,
				Address: p.Address,
			},
			Amount: p.Amount,
		}

		coins = append(coins, volume)
	}

	return coins
}
