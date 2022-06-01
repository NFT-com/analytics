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
		return bindError(fmt.Errorf("could not unpack collection request: %w", err))
	}

	// Retrieve the number of NFTs in the collection.
	count, err := a.stats.CollectionSizeHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve collection size: %w", err))
	}

	return ctx.JSON(http.StatusOK, count)
}
