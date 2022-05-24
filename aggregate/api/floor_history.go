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
		return err
	}

	// Retrieve the information about foor prices through history for the collection.
	floor, err := a.stats.CollectionFloorHistory(request.address, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not retrieve collection floor price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, floor)
}
