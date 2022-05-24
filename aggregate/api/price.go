package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NFTPrice handles the request for retrieving historic prices of an NFT.
func (a *API) NFTPrice(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup NFT identifier.
	nft, err := a.lookup.NFT(id)
	if err != nil {
		err := fmt.Errorf("could not lookup NFT: %w", err)
		return apiError(err)
	}

	// Retrieve NFT price.
	price, err := a.stats.NFTPrice(nft)
	if err != nil {
		err := fmt.Errorf("could not retrieve NFT price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, price)
}
