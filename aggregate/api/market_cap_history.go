package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MarketplaceMarketCapHistory handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCapHistory(ctx echo.Context) error {

	// Unpack and validate request
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve marketplace market cap.
	cap, err := a.stats.MarketplaceMarketCapHistory(request.addresses, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not get marketplace market cap data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}
