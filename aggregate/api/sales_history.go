package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionSalesHistory handles the request for number of sales for a collection.
func (a *API) CollectionSalesHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve number of sales for the collection.
	sales, err := a.stats.CollectionSalesHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}

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
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
