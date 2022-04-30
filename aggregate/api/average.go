package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionAverage handles the request for the average price for NFTs in a collection.
func (a *API) CollectionAverage(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	avg, err := a.stats.Average(req.ID, req.From.time(), req.To.time())
	if err != nil {
		err := fmt.Errorf("could not retrieve NFT price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, avg)
}
