package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
)

// NFTPriceHistory handles the request for retrieving historic prices of an NFT.
func (a *API) NFTPriceHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackNFTRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack NFT request: %w", err))
	}

	// Retrieve NFT prices.
	prices, err := a.stats.NFTPriceHistory(request.id, request.from, request.to)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT price history: %w", err))
	}

	return ctx.JSON(http.StatusOK, prices)
}

// NFTAveragePrice handles the request for retrieving the all-time average price of an NFT.
func (a *API) NFTAveragePrice(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup NFT identifier.
	nft, err := a.lookup.NFT(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup NFT: %w", err))
	}

	// Retrieve average price for the NFT.
	average, err := a.stats.NFTAveragePriceHistory(nft)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT average price: %w", err))
	}

	response := datapoint.Value{
		ID:    id,
		Value: average.Average,
	}

	return ctx.JSON(http.StatusOK, response)
}
