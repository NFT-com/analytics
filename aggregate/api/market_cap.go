package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionMarketCap handles the request for the market cap for a collection.
func (a *API) CollectionMarketCap(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve the collection market cap.
	cap, err := a.stats.CollectionMarketCap(address)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}

// CollectionBatchMarketCap handles the request for trading volume for a batch of collections.
func (a *API) CollectionBatchMarketCap(ctx echo.Context) error {

	var request batchRequest
	err := ctx.Bind(&request)
	if err != nil {
		return bindError(err)
	}

	// If we don't have any IDs provided, just return.
	if len(request.IDs) == 0 {
		return ctx.NoContent(http.StatusOK)
	}

	// Lookup collection addresses.
	addresses, err := a.lookup.Collections(request.IDs)
	if err != nil {
		err := fmt.Errorf("could not lookup collection addresses: %w", err)
		return apiError(err)
	}

	// Transform the map into a list of addresses.
	list := make([]identifier.Address, 0, len(addresses))
	for _, address := range addresses {
		list = append(list, address)
	}

	// Get the total volume for the collections.
	caps, err := a.stats.CollectionBatchMarketCaps(list)
	if err != nil {
		err := fmt.Errorf("could not retrieve collection market caps: %w", err)
		return apiError(err)
	}

	// Map the list of volumes back to the collection IDs.
	var marketCaps []StatValue
	for id, address := range addresses {

		cap, ok := caps[address]
		// If a collection has not been traded before, there won't be any market cap data.
		if !ok {
			a.log.Debug().Str("collection_id", id).Msg("no market cap data for collection")
			continue
		}

		// Create the volume record and add it to the list.
		v := StatValue{
			ID:    id,
			Value: cap,
		}

		marketCaps = append(marketCaps, v)
	}

	// Create the API response.
	response := BatchResponse{
		Data: marketCaps,
	}

	return ctx.JSON(http.StatusOK, response)
}

// MarketplaceMarketCap handles the request for the market cap for a marketplace.
func (a *API) MarketplaceMarketCap(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(id)
	if err != nil {
		return apiError(err)
	}

	// Retrieve marketplace market cap info.
	cap, err := a.stats.MarketplaceMarketCap(addresses)
	if err != nil {
		return apiError(err)
	}

	return ctx.JSON(http.StatusOK, cap)
}
