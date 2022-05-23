package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionVolumeHistory handles the request for the trading volume for a collection.
func (a *API) CollectionVolumeHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve collection volume.
	volume, err := a.stats.CollectionVolumeHistory(request.address, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not get volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}

// MarketplaceVolumeHistory handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolumeHistory(ctx echo.Context) error {

	// Unpack and validate request
	request, err := a.unpackMarketplaceRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve marketplace volume.
	volume, err := a.stats.MarketplaceVolumeHistory(request.addresses, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not get volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}