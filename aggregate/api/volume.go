package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionVolume handles the request for the trading volume for a collection.
func (a *API) CollectionVolume(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve collection volume.
	volume, err := a.stats.CollectionVolume(address)
	if err != nil {
		err := fmt.Errorf("could not get collection volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}

// MarketplaceVolume handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolume(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve marketplace volume info.
	volume, err := a.stats.MarketplaceVolume(addresses)
	if err != nil {
		err := fmt.Errorf("could not get marketplace volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}
