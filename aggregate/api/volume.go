package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionVolumeHistory handles the request for the trading volume for a collection.
func (a *API) CollectionVolumeHistory(ctx echo.Context) error {

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

	// Retrieve collection volume.
	volume, err := a.stats.CollectionVolumeHistory(address, req.From.time(), req.To.time())
	if err != nil {
		err := fmt.Errorf("could not get volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}

// MarketplaceVolumeHistory handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolumeHistory(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Lookup chain ID and contract addresses for the marketplace.
	addresses, err := a.lookupMarketplace(req.ID)
	if err != nil {
		return apiError(err)
	}

	// Retrieve marketplace volume.
	volume, err := a.stats.MarketplaceVolumeHistory(addresses, req.From.time(), req.To.time())
	if err != nil {
		err := fmt.Errorf("could not get volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}
