package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionCountHistory handles the request for number of NFTs in a collection.
func (a *API) CollectionCountHistory(ctx echo.Context) error {

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

	// Retrieve the number of NFTs in the collection.
	count, err := a.stats.CollectionCount(address, req.From.time(), req.To.time())
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT count: %w", err))
	}

	return ctx.JSON(http.StatusOK, count)
}
