package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MarketplaceUsersHistory handles the request for number of active users on a marketplace.
func (a *API) MarketplaceUsersHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve marketplace user data.
	users, err := a.stats.MarketplaceUserCountHistory(request.addresses, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not retrieve marketplace user count history: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, users)
}
