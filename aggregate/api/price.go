package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NFTPrice handles the request for retrieving historic prices of an NFT.
func (a *API) NFTPrice(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackNFTRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve NFT price.
	price, err := a.stats.NFTPrice(request.id)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, price)
}
