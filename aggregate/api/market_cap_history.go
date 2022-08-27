package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CollectionMarketCapHistory handles the request for the market cap for a collection.
func (a *API) CollectionMarketCapHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackCollectionHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack collection request: %w", err))
	}

	// Retrieve the collection market cap.
	mcap, err := a.stats.CollectionMarketCapHistory(request.address, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not get collection market cap data: %w", err))
	}

	// Create the API response.
	res, err := a.createValueHistoryRecord(ctx.Param(idParam), mcap)
	if err != nil {
		return apiError(fmt.Errorf("could not create response: %w", err))
	}

	return ctx.JSON(http.StatusOK, res)
}

// MarketplaceMarketCapHistory handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCapHistory(ctx echo.Context) error {

	// Unpack and validate request
	request, err := a.unpackMarketplaceHistoryRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack marketplace request: %w", err))
	}

	// Retrieve marketplace market cap.
	mcap, err := a.stats.MarketplaceMarketCapHistory(request.addresses, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not get marketplace market cap data: %w", err))
	}

	// Create the API response.
	res, err := a.createValueHistoryRecord(ctx.Param(idParam), mcap)
	if err != nil {
		return apiError(fmt.Errorf("could not create response: %w", err))
	}

	return ctx.JSON(http.StatusOK, res)
}
