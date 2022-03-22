package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Mint returns all NFT mint events, according to the specified search criteria.
func (a *API) Mint(ctx echo.Context) error {

	var req Filter
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	mints, err := a.storage.Mints(req)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, mints)
}
