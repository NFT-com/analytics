package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// CollectionVolume handles the request for the trading volume for a collection.
func (a *API) CollectionVolume(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	return a.volume(ctx, req.ID, "", req.From.time(), req.To.time())
}

// MarketplaceVolume handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolume(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	return a.volume(ctx, "", req.ID, req.From.time(), req.To.time())
}

func (a *API) volume(ctx echo.Context, collectionID string, marketplaceID string, from time.Time, to time.Time) error {

	volume, err := a.stats.Volume(collectionID, marketplaceID, from, to)
	if err != nil {
		err := fmt.Errorf("could not get volume data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, volume)
}
