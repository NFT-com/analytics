package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionAverageHistory handles the request for the average price for NFTs in a collection.
func (a *API) CollectionAverageHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve collection average value.
	avg, err := a.stats.CollectionAverageHistory(request.address, request.from, request.to)
	if err != nil {
		err := fmt.Errorf("could not retrieve collection average price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, avg)
}
