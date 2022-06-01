package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionAverageHistory handles the request for the average price for NFTs in a collection.
func (a *API) CollectionAverageHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack collection request: %w", err))
	}

	// Retrieve collection average value.
	avg, err := a.stats.CollectionAverageHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve collection average price: %w", err))
	}

	return ctx.JSON(http.StatusOK, avg)
}
