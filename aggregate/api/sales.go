package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionSales handles the request for number of sales for a collection.
func (a *API) CollectionSales(ctx echo.Context) error {

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
	sales, err := a.stats.CollectionSales(address, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}

// MarketplaceSales handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSales(ctx echo.Context) error {

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
	sales, err := a.stats.MarketplaceSales(addresses, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
