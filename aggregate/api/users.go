package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MarketplaceUsers handles the request for number of active users on a marketplace.
func (a *API) MarketplaceUsers(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackMarketplaceRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve marketplace user data.
	users, err := a.stats.MarketplaceUserCount(request.addresses)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, users)
}
