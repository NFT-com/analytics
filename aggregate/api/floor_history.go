package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionFloorHistory handles the request for floor price for a collection.
func (a *API) CollectionFloorHistory(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Lookup chain ID and contract address for the collection.
	address, err := a.lookupCollection(req.ID)
	if err != nil {
		return apiError(err)
	}

	floor, err := a.stats.CollectionFloorHistory(address, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, floor)
}
