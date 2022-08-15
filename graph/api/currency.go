package api

import (
	"fmt"

	aggregate "github.com/NFT-com/analytics/aggregate/models/api"
	graph "github.com/NFT-com/analytics/graph/models/api"
)

// FIXME: Get symbols for all currencies at once, don't do it on the fly.
func (s *Server) createCurrencyList(coins []aggregate.Coin) ([]graph.Currency, error) {

	out := make([]graph.Currency, 0, len(coins))
	for _, coin := range coins {

		symbol, err := s.storage.CurrencySymbol(coin.CurrencyID)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve currency symbol: %w", err)
		}

		currency := graph.Currency{
			Amount: coin.Amount,
			Symbol: symbol,
		}

		out = append(out, currency)
	}

	return out, nil
}
