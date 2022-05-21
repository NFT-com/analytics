package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionSalesHistory handles the request for number of sales for a collection.
func (a *API) CollectionSalesHistory(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Lookup chain ID and contract address for the collection.
	address, err := a.lookupCollection(req.ID)
	if err != nil {
		return apiError(err)
	}

	// Retrieve number of sales for the collection.
	sales, err := a.stats.CollectionSalesHistory(address, req.From, req.To)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}

// MarketplaceSalesHistory handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSalesHistory(ctx echo.Context) error {

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

	// Retrieve number of sales for the collection.
	sales, err := a.stats.MarketplaceSalesHistory(addresses, req.From, req.To)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
