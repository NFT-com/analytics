package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionFloorHistory handles the request for floor price for a collection.
func (a *API) CollectionFloorHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack collection request: %w", err))
	}

	// Retrieve the information about foor prices through history for the collection.
	floor, err := a.stats.CollectionFloorHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve collection floor price: %w", err))
	}

	return ctx.JSON(http.StatusOK, floor)
}
