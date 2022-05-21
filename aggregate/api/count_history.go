package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionCountHistory handles the request for number of NFTs in a collection.
func (a *API) CollectionCountHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve the number of NFTs in the collection.
	count, err := a.stats.CollectionCountHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT count: %w", err))
	}

	return ctx.JSON(http.StatusOK, count)
}
