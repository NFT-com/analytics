package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionFloorHistory handles the request for floor price for a collection.
func (a *API) CollectionFloorHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionRequest(ctx)
	if err != nil {
		return err
	}

	floor, err := a.stats.CollectionFloorHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, floor)
}
