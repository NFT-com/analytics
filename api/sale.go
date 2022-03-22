package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Sale returns all NFT sale events, according to the specified search criteria.
func (a *API) Sale(ctx echo.Context) error {

	var req Filter
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	sales, err := a.storage.Sales(req)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, sales)
}
