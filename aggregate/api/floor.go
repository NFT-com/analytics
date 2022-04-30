package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionFloor handles the request for floor price for a collection.
func (a *API) CollectionFloor(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	floor, err := a.stats.Floor(req.ID, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, floor)
}
