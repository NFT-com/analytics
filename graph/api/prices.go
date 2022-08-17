package api

import (
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/NFT-com/analytics/graph/api/internal/query"
	"github.com/NFT-com/analytics/graph/models/api"
)

func (s *Server) getNFTStats(query *query.NFT, nft *api.NFT) error {

	var multiErr error

	// Retrieve NFT price from the aggregation API.
	if query.Price {
		price, err := s.aggregationAPI.Price(nft.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not retrieve price for NFT: %w", err))
		} else {

			// Translate the Aggregation API format to the expected Graph format.
			formatted, err := s.convertCoinsToCurrencies(price)
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf("could not convert price coin list to currencies: %w", err))
			}

			nft.TradingPrice = formatted
		}
	}

	// Retrieve NFT average price from the aggregation API.
	if query.AveragePrice {
		average, err := s.aggregationAPI.AveragePrice(nft.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not retrieve average price for NFT: %w", err))
		} else {

			// Translate the Aggregation API format to the expected Graph format.
			formatted, err := s.convertCoinsToCurrencies(average)
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf("could not convert average price coin list to currencies: %w", err))
			}

			nft.AveragePrice = formatted
		}
	}

	return multiErr
}
