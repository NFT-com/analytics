package api

import (
	"fmt"
	"net/http"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/labstack/echo/v4"
)

// CollectionSales handles the request for number of sales for a collection.
func (a *API) CollectionSales(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(err)
	}
	// Retrieve number of sales for the collection.
	sales, err := a.stats.CollectionSales(address)
	if err != nil {
		err := fmt.Errorf("could not retrieve collection sales: %w", err)
		return apiError(err)
	}

	response := datapoint.Value{
		ID:    id,
		Value: float64(sales),
	}

	return ctx.JSON(http.StatusOK, response)
}

// MarketplaceSales handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSales(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve number of sales for the marketplace.
	sales, err := a.stats.MarketplaceSales(addresses)
	if err != nil {
		err := fmt.Errorf("could not retrieve marketplace sales: %w", err)
		return apiError(err)
	}

	response := datapoint.Value{
		ID:    id,
		Value: float64(sales),
	}

	return ctx.JSON(http.StatusOK, response)
}
