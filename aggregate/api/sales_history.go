package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MarketplaceSalesHistory handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSalesHistory(ctx echo.Context) error {

	// Unpack and validate request
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve number of sales for the collection.
	sales, err := a.stats.MarketplaceSalesHistory(request.addresses, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not retrieve marketplace sales history: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
