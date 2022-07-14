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
		}

		marketplace.Volume, _ = volume.Float64()
	}

	// Get market cap from the aggregation API.
	if query.MarketCap {
		cap, err := s.aggregationAPI.MarketplaceMarketCap(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace market cap: %w", err))
		}

		marketplace.MarketCap, _ = cap.Float64()
	}

	// Get sale count from the aggregation API.
	if query.Sales {
		sales, err := s.aggregationAPI.MarketplaceSales(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace sales: %w", err))
		}

		marketplace.Sales = sales.BigInt().Uint64()
	}

	// Get user count from the aggregation API.
	if query.Users {
		users, err := s.aggregationAPI.MarketplaceUsers(marketplace.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get marketplace users: %w", err))
		}

		marketplace.Users = users.BigInt().Uint64()
	}

	return multiErr
}
