package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionMarketCap handles the request for the market cap for a collection.
func (a *API) CollectionMarketCap(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve the collection market cap.
	cap, err := a.stats.CollectionMarketCap(address)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}

// MarketplaceMarketCap handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCap(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve marketplace market cap info.
	cap, err := a.stats.MarketplaceMarketCap(addresses)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}
