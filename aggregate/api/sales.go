package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// CollectionSales handles the request for number of sales for a collection.
func (a *API) CollectionSales(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	return a.sales(ctx, req.ID, "", req.From.time(), req.To.time())
}

// MarketplaceSales handles the request for number of sales for a marketplace.
func (a *API) MarketplaceSales(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	return a.sales(ctx, "", req.ID, req.From.time(), req.To.time())
}

func (a *API) sales(ctx echo.Context, collectionID string, marketplaceID string, from time.Time, to time.Time) error {

	sales, err := a.stats.Sales(collectionID, marketplaceID, from, to)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
