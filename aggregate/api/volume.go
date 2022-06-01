package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/aggregate/models/api"
	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// CollectionVolume handles the request for the trading volume for a collection.
func (a *API) CollectionVolume(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup collection address.
	address, err := a.lookupCollection(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup collection: %w", err))
	}

	// Retrieve collection volume.
	volume, err := a.stats.CollectionVolume(address)
	if err != nil {
		return apiError(fmt.Errorf("could not get collection volume data: %w", err))
	}

	response := datapoint.Value{
		ID:    id,
		Value: volume,
	}

	return ctx.JSON(http.StatusOK, response)
}

// CollectionBatchVolume handles the request for trading volume for a batch of collections.
func (a *API) CollectionBatchVolume(ctx echo.Context) error {

	var request api.BatchRequest
	err := ctx.Bind(&request)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack collection batch request: %w", err))
	}

	// If we don't have any IDs provided, just return.
	if len(request.IDs) == 0 {
		return ctx.NoContent(http.StatusOK)
	}

	// Lookup collection addresses.
	addresses, err := a.lookupCollections(request.IDs)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup collection addresses: %w", err))
	}

	// Transform the map into a list of addresses.
	list := make([]identifier.Address, 0, len(addresses))
	for _, address := range addresses {
		list = append(list, address)
	}

	// Get the total volume for the collections.
	volumes, err := a.stats.CollectionBatchVolumes(list)
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve collection volumes: %w", err))
	}

	// Map the list of volumes back to the collection IDs.
	var collectionVolumes []datapoint.Value
	for id, address := range addresses {

		volume, ok := volumes[address]
		// If a collection has not been traded before, there won't be any volume data.
		if !ok {
			a.log.Debug().Str("collection_id", id).Msg("no volume data for collection")
			continue
		}

		// Create the volume record and add it to the list.
		v := datapoint.Value{
			ID:    id,
			Value: volume,
		}

		collectionVolumes = append(collectionVolumes, v)
	}

	// Create the API response.
	response := api.BatchResponse{
		Data: collectionVolumes,
	}

	return ctx.JSON(http.StatusOK, response)
}

// MarketplaceVolume handles the request for the trading volume for a marketplace.
func (a *API) MarketplaceVolume(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup marketplace: %w", err))
	}

	// Retrieve marketplace volume info.
	volume, err := a.stats.MarketplaceVolume(addresses)
	if err != nil {
		return apiError(fmt.Errorf("could not get marketplace volume data: %w", err))
	}

	response := datapoint.Value{
		ID:    id,
		Value: volume,
	}

	return ctx.JSON(http.StatusOK, response)
}
