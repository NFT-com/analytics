package api

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

func (a *API) unpackCollectionHistoryRequest(ctx echo.Context) (*collectionRequest, error) {

	// Unpack request.
	var request apiRequest
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
		from:    request.From.time(),
		to:      request.To.time(),
	}

	return out, nil
}

func (a *API) unpackMarketplaceHistoryRequest(ctx echo.Context) (*marketplaceRequest, error) {

	// Unpack the request.
	var request apiRequest
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
		from:      request.From.time(),
		to:        request.To.time(),
	}

	return out, nil
}

func (a *API) unpackNFTRequest(ctx echo.Context) (*nftRequest, error) {

	// Unpack request.
	var request apiRequest
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
		from: request.From.time(),
		to:   request.To.time(),
	}

	return out, nil
}

func (a *API) lookupCollection(id string) (identifier.Address, error) {

	address, ok := a.collectionCache[id]
	if ok {
		return address, nil
	}

	address, err := a.lookup.Collection(id)
	if err != nil {
		return identifier.Address{}, fmt.Errorf("could not lookup collection: %w", err)
	}

	// FIXME: Add a mutex to sync this.
	a.collectionCache[id] = address

	return address, nil
}

func (a *API) lookupMarketplace(id string) ([]identifier.Address, error) {

	addresses, err := a.lookup.Marketplace(id)
	if err != nil {
		return nil, fmt.Errorf("could not lookup marketplace: %w", err)
	}

	// FIXME: Add caching.

	return addresses, nil
}
