package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// MarketplaceUsersHistory handles the request for number of active users on a marketplace.
func (a *API) MarketplaceUsersHistory(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Lookup chain ID and contract address for the marketplace.
	addresses, err := a.lookupMarketplace(req.ID)
	if err != nil {
		return apiError(err)
	}

	users, err := a.stats.MarketplaceUserCount(addresses, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, users)
}
