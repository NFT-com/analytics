package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionSales handles the request for number of sales for a collection.
func (a *API) CollectionSales(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve number of sales for the collection.
	sales, err := a.stats.CollectionSales(request.address)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}

// MarketplaceSales handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSales(ctx echo.Context) error {

	// Unpack and validate request
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve number of sales for the marketplace.
	sales, err := a.stats.MarketplaceSales(request.addresses)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
