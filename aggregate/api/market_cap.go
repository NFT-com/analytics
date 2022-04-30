package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// CollectionMarketCap handles the request for the market cap for a collection.
func (a *API) CollectionMarketCap(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	return a.marketCap(ctx, req.ID, "", req.From.time(), req.To.time())
}

// MarketplaceMarketCap handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCap(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	return a.marketCap(ctx, "", req.ID, req.From.time(), req.To.time())
}

func (a *API) marketCap(ctx echo.Context, collectionID string, marketplaceID string, from time.Time, to time.Time) error {

	cap, err := a.stats.MarketCap(collectionID, marketplaceID, from, to)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}
