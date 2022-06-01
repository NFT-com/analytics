package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionSalesHistory handles the request for number of sales for a collection.
func (a *API) CollectionSalesHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack collection request: %w", err))
	}

	// Retrieve number of sales for the collection.
	sales, err := a.stats.CollectionSalesHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve collection sales history: %w", err))
	}

	return ctx.JSON(http.StatusOK, sales)
}

// MarketplaceSalesHistory handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSalesHistory(ctx echo.Context) error {

	// Unpack and validate request
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack marketplace request: %w", err))
	}

	// Retrieve number of sales for the collection.
	sales, err := a.stats.MarketplaceSalesHistory(request.addresses, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve marketplace sales history: %w", err))
	}

	return ctx.JSON(http.StatusOK, sales)
}
