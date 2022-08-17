package api

import (
	"fmt"

	aggregate "github.com/NFT-com/analytics/aggregate/models/api"
	graph "github.com/NFT-com/analytics/graph/models/api"
)

// TODO: Consider retrieving all currency symbols on startup, then just retrieving them from a map when needed.
// Seehttps://github.com/NFT-com/analytics/issues/87
func (s *Server) convertCoinsToCurrencies(coins []aggregate.Coin) ([]graph.Currency, error) {

	out := make([]graph.Currency, 0, len(coins))
	for _, coin := range coins {

		symbol, err := s.storage.CurrencySymbol(coin.CurrencyID)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve currency symbol: %w", err)
		}

		currency := graph.Currency{
			Symbol: symbol,
			Value:  coin.Value,
		}

		out = append(out, currency)
	}

	return out, nil
}
