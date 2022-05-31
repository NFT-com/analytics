package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionSizeHistory handles the request for number of NFTs in a collection.
func (a *API) CollectionSizeHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve the number of NFTs in the collection.
	count, err := a.stats.CollectionSizeHistory(request.address, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not retrieve collection size: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, count)
}
