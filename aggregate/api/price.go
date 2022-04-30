package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NFTPrice handles the request for retrieving historic prices of an NFT.
func (a *API) NFTPrice(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	price, err := a.stats.NFTPrice(req.ID, req.From.time(), req.To.time())
	if err != nil {
		err := fmt.Errorf("could not retrieve NFT price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, price)
}

// NFTAveragePrice handles the request for retrieving the all-time average price of an NFT.
func (a *API) NFTAveragePrice(ctx echo.Context) error {

	id := ctx.Param(idParamName)

	avg, err := a.stats.NFTAveragePrice(id)
	if err != nil {
		err := fmt.Errorf("could not retrieve average price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, avg)
}
