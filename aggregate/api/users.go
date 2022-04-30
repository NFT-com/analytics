package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MarketplaceUsers handles the request for number of active users on a marketplace.
func (a *API) MarketplaceUsers(ctx echo.Context) error {

	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	log.Printf("marketplace users request: %+#v, from: %v, to: %v", req,
		req.From.time().Format(timeFormat),
		req.To.time().Format(timeFormat),
	)

	count, err := a.stats.MarketplaceUsers(req.ID, req.From.time(), req.To.time())
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT count: %w", err))
	}

	return ctx.JSON(http.StatusOK, count)
}
