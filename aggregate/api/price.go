package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
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

	response := api.Value{
		ID:    id,
		Value: price,
	}

	return ctx.JSON(http.StatusOK, response)
}

// CollectionPrices handles the request for retrieving current prices for NFTs in a collection.
func (a *API) CollectionPrices(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup collection: %w", err))
	}

	// Retrieve NFT IDs in order to later link prices to the NFT identifiers.
	nftIDs, err := a.lookup.CollectionNFTs(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup collection NFTs: %w", err))
	}

	// Retrieve NFT prices.
	prices, err := a.stats.CollectionPrices(address)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT prices: %w", err))
	}

	// Link retrieved prices to the NFT by ID.
	var nftPrices []api.Values
	for id, nftAddress := range nftIDs {

		price, ok := prices[lowerNFTID(nftAddress)]
		// If an NFT has never been sold before, it's normal that we don't know the price.
		if !ok {
			a.log.Debug().Str("nft_id", id).Msg("no known price for NFT")
			continue
		}

		// FIXME: Think about this again.
		// Create the price record and add it to the list.
		p := api.Values{
			ID:     id,
			Values: []datapoint.Currency{price},
		}
		nftPrices = append(nftPrices, p)
	}

	// Create the API response.
	response := api.BatchResponse{
		Data: nftPrices,
	}

	return ctx.JSON(http.StatusOK, response)
}

// CollectionAveragePrices handles the request for retrieving all-time average prices for NFTs in a collection.
func (a *API) CollectionAveragePrices(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup collection: %w", err))
	}

	// Retrieve NFT IDs in order to later link prices to the NFT identifiers.
	nftIDs, err := a.lookup.CollectionNFTs(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup collection NFTs: %w", err))
	}

	// Retrieve NFT average averages.
	averages, err := a.stats.CollectionAveragePrices(address)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT average prices: %w", err))
	}

	// Link retrieved prices to the NFT by ID.
	var nftPrices []api.Values
	for id, nftAddress := range nftIDs {

		average, ok := averages[lowerNFTID(nftAddress)]
		// If an NFT has never been sold before, it's normal that there's no average price.
		if !ok {
			a.log.Debug().Str("nft_id", id).Msg("no average price for NFT")
			continue
		}

		// Create the price record and add it to the list.
		p := api.Values{
			ID:     id,
			Values: average,
		}
		nftPrices = append(nftPrices, p)
	}

	// Create the API response.
	response := api.BatchResponse{
		Data: nftPrices,
	}

	return ctx.JSON(http.StatusOK, response)
}
