package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionAverageHistory handles the request for the average price for NFTs in a collection.
func (a *API) CollectionAverageHistory(ctx echo.Context) error {

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

	// Retrieve collection average value.
	avg, err := a.stats.CollectionAverage(address, req.From.time(), req.To.time())
	if err != nil {
		err := fmt.Errorf("could not retrieve collection average price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, avg)
}
