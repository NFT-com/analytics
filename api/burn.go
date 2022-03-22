package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Burn returns all NFT burn events, according to the specified search criteria.
func (a *API) Burn(ctx echo.Context) error {

	var req Filter
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	burns, err := a.storage.Burns(req)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, burns)
}
