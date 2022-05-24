package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionVolume handles the request for the trading volume for a collection.
func (a *API) CollectionVolume(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve collection volume.
	volume, err := a.stats.CollectionVolume(request.address)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}

// MarketplaceVolume handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolume(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve marketplace volume info.
	volume, err := a.stats.MarketplaceVolume(request.addresses)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}
