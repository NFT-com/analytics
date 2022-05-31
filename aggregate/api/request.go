package api

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/aggregate/models/api"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

func (a *API) unpackCollectionHistoryRequest(ctx echo.Context) (*collectionRequest, error) {

	// Unpack request.
	var request api.Request
	err := ctx.Bind(&request)
	if err != nil {
		return nil, bindError(err)
	}

	if request.ID == "" {
		err = errors.New("collection ID is required")
		return nil, bindError(err)
	}

	// Lookup collection address.
	address, err := a.lookupCollection(request.ID)
	if err != nil {
		return nil, apiError(err)
	}

	out := &collectionRequest{
		address: address,
		from:    request.From.Time(),
		to:      request.To.Time(),
	}

	return out, nil
}

func (a *API) unpackMarketplaceHistoryRequest(ctx echo.Context) (*marketplaceRequest, error) {

	// Unpack the request.
	var request api.Request
	err := ctx.Bind(&request)
	if err != nil {
		return nil, bindError(err)
	}

	if request.ID == "" {
		err = errors.New("marketplace ID is required")
		return nil, bindError(err)
	}

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(request.ID)
	if err != nil {
		return nil, apiError(err)
	}

	out := &marketplaceRequest{
		addresses: addresses,
		from:      request.From.Time(),
		to:        request.To.Time(),
	}

	return out, nil
}

func (a *API) unpackNFTRequest(ctx echo.Context) (*nftRequest, error) {

	// Unpack request.
	var request api.Request
	err := ctx.Bind(&request)
	if err != nil {
		return nil, bindError(err)
	}

	if request.ID == "" {
		err = errors.New("NFT ID is required")
		return nil, bindError(err)
	}

	// Lookup NFT identifier.
	nft, err := a.lookup.NFT(request.ID)
	if err != nil {
		return nil, fmt.Errorf("could not lookup NFT: %w", err)
	}

	out := &nftRequest{
		id:   nft,
		from: request.From.Time(),
		to:   request.To.Time(),
	}

	return out, nil
}

// lookupCollections is a helper function wrapping `lookupCollection`, operating on a list of collection IDs.
func (a *API) lookupCollections(ids []string) (map[string]identifier.Address, error) {

	addresses := make(map[string]identifier.Address, len(ids))
	for _, id := range ids {

		address, err := a.lookupCollection(id)
		if err != nil {
			return nil, fmt.Errorf("could not lookup collection: %w", err)
		}

		addresses[id] = address
	}

	return addresses, nil
}

func (a *API) lookupCollection(id string) (identifier.Address, error) {

	addresses, ok := a.collections.get(id)
	if ok {
		// Just a safety check, we should never have more than one address for a collection ID.
		if len(addresses) != 1 {
			return identifier.Address{}, fmt.Errorf("unexpected number of collection addresses (have: %d)", len(addresses))
		}

		return addresses[0], nil
	}

	address, err := a.lookup.Collection(id)
	if err != nil {
		return identifier.Address{}, fmt.Errorf("could not lookup collection: %w", err)
	}

	a.collections.set(id, []identifier.Address{address})

	return address, nil
}

func (a *API) lookupMarketplace(id string) ([]identifier.Address, error) {

	addresses, ok := a.marketplaces.get(id)
	if ok {
		return addresses, nil
	}

	addresses, err := a.lookup.Marketplace(id)
	if err != nil {
		return nil, fmt.Errorf("could not lookup marketplace: %w", err)
	}

	a.marketplaces.set(id, addresses)

	return addresses, nil
}
