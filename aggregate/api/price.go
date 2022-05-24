package api

import (
	"fmt"
	"net/http"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
	"github.com/labstack/echo/v4"
)

// NFTPrice handles the request for retrieving current price of an NFT.
func (a *API) NFTPrice(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup NFT identifier.
	nft, err := a.lookup.NFT(id)
	if err != nil {
		err := fmt.Errorf("could not lookup NFT: %w", err)
		return apiError(err)
	}

	// Retrieve NFT price.
	price, err := a.stats.NFTPrice(nft)
	if err != nil {
		err := fmt.Errorf("could not retrieve NFT price: %w", err)
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, price)
}

// NFTBatchPrice handles the request for retrieving current price for a batch of NFTs.
func (a *API) NFTBatchPrice(ctx echo.Context) error {

	// Unpack the list of IDs.
	var request batchRequest
	err := ctx.Bind(&request)
	if err != nil {
		return bindError(err)
	}

	// If we don't have any IDs provided, just return.
	if len(request.IDs) == 0 {
		return ctx.NoContent(http.StatusNoContent)
	}

	// Lookup the list of NFT identifiers based on IDs.
	addresses, err := a.lookup.NFTs(request.IDs)
	if err != nil {
		err := fmt.Errorf("could not lookup NFT addresses: %w", err)
		return apiError(err)
	}

	// Transform the map into a list of identifiers.
	list := make([]identifier.NFT, 0, len(addresses))
	for _, address := range addresses {
		list = append(list, address)
	}

	// Get the prices for the NFT set.
	prices, err := a.stats.NFTBatchPrices(list)
	if err != nil {
		err := fmt.Errorf("could not retrieve batch prices: %w", err)
		return apiError(err)
	}

	// Map the list of prices back to the NFT IDs.
	var nftPrices []NFTPrice
	for id, address := range addresses {

		price, ok := prices[address]

		// If an NFT has never been sold before, it's normal that we don't know the price.
		if !ok {
			a.log.Debug().Str("nft_id", id).Msg("no known price for NFT")
			continue
		}

		// Create the price record and add it to the list.

		p := NFTPrice{
			ID:    id,
			Price: price.Price,
		}

		nftPrices = append(nftPrices, p)
	}

	// Create the API response.
	response := NFTPriceResponse{
		Prices: nftPrices,
	}

	return ctx.JSON(http.StatusOK, response)
}
