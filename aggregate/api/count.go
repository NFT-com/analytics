package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionCount handles the request for number of NFTs in a collection.
func (a *API) CollectionCount(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	count, err := a.stats.Count(req.ID, req.From.time(), req.To.time())
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT count: %w", err))
	}

	return ctx.JSON(http.StatusOK, count)
}
