package api

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

func (a *API) unpackCollectionRequest(ctx echo.Context) (*collectionRequest, error) {

	// Unpack request.
	var request apiRequest
	err := ctx.Bind(&request)
	if err != nil {
		return nil, bindError(err)
	}

	// Validate request data.
	err = a.validate.Struct(&request)
	if err != nil {
		return nil, bindError(err)
	}

	// Lookup collection address.
	address, err := a.lookupCollection(request.ID)
	if err != nil {
		return nil, apiError(err)
	}

	out := &collectionRequest{
		address: address,
		from:    request.From,
		to:      request.To,
	}

	return out, nil
}

func (a *API) unpackMarketplaceRequest(ctx echo.Context) (*marketplaceRequest, error) {

	// Unpack the request.
	var request apiRequest
	err := ctx.Bind(&request)
	if err != nil {
		return nil, bindError(err)
	}

	// Validate the request data.
	err = a.validate.Struct(&request)
	if err != nil {
		return nil, bindError(err)
	}

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(request.ID)
	if err != nil {
		return nil, apiError(err)
	}

	out := &marketplaceRequest{
		addresses: addresses,
		from:      request.From,
		to:        request.To,
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

func (a *API) lookupNFT(id string) (identifier.NFT, error) {

	nft, err := a.lookup.NFT(id)
	if err != nil {
		return identifier.NFT{}, fmt.Errorf("could not lookup NFT: %w", err)
	}

	// FIXME: Add caching. Though, due to the amount of NFTs, it may not be feasible
	// to cache ALL NFT IDs. Instad, perhaps handle them by group.

	return nft, nil
}

func (a *API) lookupMarketplace(id string) ([]identifier.Address, error) {

	addresses, err := a.lookup.Marketplace(id)
	if err != nil {
		return nil, fmt.Errorf("could not lookup marketplace: %w", err)
	}

	// FIXME: Add caching.

	return addresses, nil
}
