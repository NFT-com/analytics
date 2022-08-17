package api

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

func (a *API) unpackCollectionHistoryRequest(ctx echo.Context) (*collectionRequest, error) {

	// Unpack request.
	var request api.Request
	err := ctx.Bind(&request)
	if err != nil {
		return nil, bindError(err)
	}

	if request.ID == "" {
		return nil, errors.New("collection ID is required")
	}

	// Lookup collection address.
	address, err := a.lookupCollection(request.ID)
	if err != nil {
		return nil, fmt.Errorf("could not lookup collection: %w", err)
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
		return nil, errors.New("marketplace ID is required")
	}

	// Lookup marketplace addresses.
	addresses, err := a.lookupMarketplace(request.ID)
	if err != nil {
		return nil, fmt.Errorf("could not lookup marketplace: %w", err)
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
		return nil, errors.New("NFT ID is required")
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

	addresses, ok := a.collections.Get(id)
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

	a.collections.Set(id, []identifier.Address{address})

	return address, nil
}

func (a *API) lookupMarketplace(id string) ([]identifier.Address, error) {

	addresses, ok := a.marketplaces.Get(id)
	if ok {
		return addresses, nil
	}

	addresses, err := a.lookup.Marketplace(id)
	if err != nil {
		return nil, fmt.Errorf("could not lookup marketplace: %w", err)
	}

	a.marketplaces.Set(id, addresses)

	return addresses, nil
}

func (a *API) lookupCurrencyID(currency identifier.Currency) (string, error) {

	id, ok := a.currencies.Get(currency)
	if ok {
		return id, nil
	}

	id, err := a.lookup.CurrencyID(currency)
	if err != nil {
		return "", fmt.Errorf("could not lookup currency: %w", err)
	}

	a.currencies.Set(currency, id)

	return id, nil
}
