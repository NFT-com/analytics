package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/NFT-com/analytics/aggregate/models/api"
)

// NFTPriceHistory handles the request for retrieving historic prices of an NFT.
func (a *API) NFTPriceHistory(ctx echo.Context) error {

	// Unpack and validate request.
	request, err := a.unpackNFTRequest(ctx)
	if err != nil {
		return bindError(fmt.Errorf("could not unpack NFT request: %w", err))
	}

	response := api.ValueHistory{
		ID:        ctx.Param(idParam),
		Snapshots: []api.CoinSnapshot{},
	}

	// Retrieve NFT prices.
	prices, err := a.stats.NFTPriceHistory(request.id, request.from, request.to)
	if err != nil && errors.Is(err, ErrRecordNotFound) {
		// The NFT has no sales.
		return ctx.JSON(http.StatusOK, response)
	}
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT price history: %w", err))
	}

	// Translate to the API data format.
	coins := make([]api.CoinSnapshot, 0, len(prices))
	for _, p := range prices {

		if p.Coin.Currency.Address == "" {
			continue
		}

		id, err := a.lookupCurrencyID(p.Coin.Currency)
		if err != nil {
			return apiError(fmt.Errorf("could not lookup currency ID: %w", err))
		}

		coin := api.Coin{
			CurrencyID: id,
			Value:      p.Coin.Value,
		}

		snapshot := api.CoinSnapshot{
			Value: []api.Coin{coin},
			Time:  p.Time,
		}

		coins = append(coins, snapshot)
	}

	response.Snapshots = coins

	return ctx.JSON(http.StatusOK, response)
}

// NFTAveragePrice handles the request for retrieving the all-time average price of an NFT.
func (a *API) NFTAveragePrice(ctx echo.Context) error {

	id := ctx.Param(idParam)

	// Lookup NFT identifier.
	nft, err := a.lookup.NFT(id)
	if err != nil {
		return apiError(fmt.Errorf("could not lookup NFT: %w", err))
	}

	response := api.Value{
		ID:    id,
		Value: []api.Coin{},
	}

	// Retrieve average price for the NFT.
	average, err := a.stats.NFTAveragePrice(nft)
	if err != nil && errors.Is(err, ErrRecordNotFound) {
		// The NFT has no sales.
		return ctx.JSON(http.StatusOK, response)
	}
	if err != nil {
		return apiError(fmt.Errorf("could not retrieve NFT average price: %w", err))
	}

	// Translate the datapoint Coin format to the API format.
	value, err := a.createCoinList(average)
	if err != nil {
		return apiError(fmt.Errorf("could not create coin list: %w", err))
	}

	response.Value = value

	return ctx.JSON(http.StatusOK, response)
}
