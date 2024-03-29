package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
)

// MarketplaceUsers handles the request for number of active users on a marketplace.
func (a *API) MarketplaceUsers(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup marketplace: %w", err))
	}

	// Retrieve marketplace user data.
	users, err := a.stats.MarketplaceUserCount(addresses)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve marketplace user count: %w", err))
	}

	response := datapoint.Count{
		ID:    id,
		Value: users,
	}

	return ctx.JSON(http.StatusOK, response)
}
