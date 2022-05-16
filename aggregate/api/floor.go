package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionFloor handles the request for floor price for a collection.
func (a *API) CollectionFloor(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Lookup chain ID and contract address for the collection.
	chainID, address, err := a.lookupCollection(req.ID)
	if err != nil {
		return apiError(err)
	}

	floor, err := a.stats.CollectionFloor(chainID, address, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, floor)
}
