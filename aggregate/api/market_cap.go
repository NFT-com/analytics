package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionMarketCap handles the request for the market cap for a collection.
func (a *API) CollectionMarketCap(ctx echo.Context) error {

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

	// Retrieve the collection market cap.
	cap, err := a.stats.CollectionMarketCap(address, req.From.time(), req.To.time())
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}

// MarketplaceMarketCap handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCap(ctx echo.Context) error {

	// Unpack the request.
	var req apiRequest
	err := ctx.Bind(&req)
	if err != nil {
		return bindError(err)
	}

	// Lookup chain ID and contract addresses for the marketplace.
	addresses, err := a.lookupMarketplace(req.ID)
	if err != nil {
		return apiError(err)
	}

	// Retrieve marketplace market cap.
	cap, err := a.stats.MarketplaceMarketCap(addresses, req.From.time(), req.To.time())
	if err != nil {
		err := fmt.Errorf("could not get market cap data: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}
