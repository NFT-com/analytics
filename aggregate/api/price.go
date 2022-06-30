package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// NFTPrice handles the request for retrieving current price of an NFT.
func (a *API) NFTPrice(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup NFT identifier.
	nft, err := a.lookup.NFT(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup NFT: %w", err))
	}

	// Retrieve NFT price.
	price, err := a.stats.NFTPrice(nft)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT price: %w", err))
	}

	response := datapoint.Value{
		ID:    id,
		Value: price,
	}

	return ctx.JSON(http.StatusOK, response)
}

// NFTBatchPrice handles the request for retrieving current price for a batch of NFTs.
func (a *API) NFTBatchPrice(ctx echo.Context) error {

	// Unpack the list of IDs.
	var request api.BatchRequest
	err := ctx.Bind(&request)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack NFT batch price request: %w", err))
	}

	// If we don't have any IDs provided, just return.
	if len(request.IDs) == 0 {
		return ctx.NoContent(http.StatusOK)
	}

	// Lookup the list of NFT identifiers based on IDs.
	addresses, err := a.lookup.NFTs(request.IDs)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup NFT addresses: %w", err))
	}

	// Transform the map into a list of identifiers.
	list := make([]identifier.NFT, 0, len(addresses))
	for _, address := range addresses {
		list = append(list, address)
	}

	// Get the prices for the NFT set.
	prices, err := a.stats.NFTBatchPrices(list)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve batch prices: %w", err))
	}

	// Map the list of prices back to the NFT IDs.
	var nftPrices []datapoint.Value
	for id, address := range addresses {

		price, ok := prices[address]
		// If an NFT has never been sold before, it's normal that we don't know the price.
		if !ok {
			a.log.Debug().Str("nft_id", id).Msg("no known price for NFT")
			continue
		}

		// Create the price record and add it to the list.
		p := datapoint.Value{
			ID:    id,
			Value: price,
		}
		nftPrices = append(nftPrices, p)
	}

	// Create the API response.
	response := api.BatchResponse{
		Data: nftPrices,
	}

	return ctx.JSON(http.StatusOK, response)
}
