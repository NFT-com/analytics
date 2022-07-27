package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionLowestPriceHistory handles the request for getting the lowest NFT price for a collection.
func (a *API) CollectionLowestPriceHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack collection request: %w", err))
	}

	// Retrieve the information about lowest prices through history for the collection.
	lowest, err := a.stats.CollectionLowestPriceHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve collection lowest price: %w", err))
	}

	return ctx.JSON(http.StatusOK, lowest)
}
