package api

import (
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/NFT-com/analytics/graph/api/internal/query"
	"github.com/NFT-com/analytics/graph/models/api"
)

func (s *Server) expandMarketplaceStats(query *query.Marketplace, marketplace *api.Marketplace) error {

	// Execute as much as possible, return the composite error in the end.
	var multiErr error

	// Get volume from the aggregation API.
	if query.Volume {
		volume, err := s.aggregationAPI.MarketplaceVolume(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace volume: %w", err))
		} else {

			// Translate the Aggregation API format to the expected Graph format.
			formatted, err := s.convertCoinsToCurrencies(volume)
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf("could not convert volume coin list to currencies: %w", err))
			}

			marketplace.Volume = formatted
		}
	}

	// Get market cap from the aggregation API.
	if query.MarketCap {
		mcap, err := s.aggregationAPI.MarketplaceMarketCap(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace market cap: %w", err))
		} else {

			// Translate the Aggregation API format to the expected Graph format.
			formatted, err := s.convertCoinsToCurrencies(mcap)
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf("could not convert market cap coin list to currencies: %w", err))
			}

			marketplace.MarketCap = formatted
		}
	}

	// Get sale count from the aggregation API.
	if query.Sales {
		sales, err := s.aggregationAPI.MarketplaceSales(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace sales: %w", err))
		}

		marketplace.Sales = sales
	}

	// Get user count from the aggregation API.
	if query.Users {
		users, err := s.aggregationAPI.MarketplaceUsers(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace users: %w", err))
		}

		marketplace.Users = users
	}

	return multiErr
}
