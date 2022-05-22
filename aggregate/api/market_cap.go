package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionMarketCap handles the request for the market cap for a collection.
func (a *API) CollectionMarketCap(ctx echo.Context) error {

	// FIXME: Perhaps don't use the full structure for the request here?

	// Unpack and validate request.
	request, err := a.unpackCollectionRequest(ctx)
	if err != nil {
		return err
	}

	// Retrieve the collection market cap.
	cap, err := a.stats.CollectionMarketCap(request.address)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}

// MarketplaceMarketCap handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCap(ctx echo.Context) error {
	return errors.New("TBD: Not implemented")
}
