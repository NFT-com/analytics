package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NFTPrice handles the request for retrieving historic prices of an NFT.
func (a *API) NFTPrice(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Retrieve contract address and token ID.
	nft, err := a.lookupNFT(req.ID)
	if err != nil {
		return apiError(err)
	}

	// Retrieve NFT prices.
	prices, err := a.stats.NFTPrice(nft, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, prices)
}

// NFTAveragePrice handles the request for retrieving the all-time average price of an NFT.
func (a *API) NFTAveragePrice(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}
