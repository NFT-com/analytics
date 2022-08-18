package api

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/api"
	"github.com/NFT-com/analytics/aggregate/models/datapoint"
)

// createCoinList takes a list of coins and transforms them the the API data format,
// translating chain ID and currency address pairs to the Currency ID.
func (a *API) createCoinList(currencies []datapoint.Coin) ([]api.Coin, error) {

	out := make([]api.Coin, 0, len(currencies))
	for _, currency := range currencies {

		// Skip numbers for sales events with no currency information.
		if currency.Currency.Address == "" {
			continue
		}

		id, err := a.lookupCurrencyID(currency.Currency)
		if err != nil {
			return nil, fmt.Errorf("could not lookup currency ID (chain: %d, address: %s): %w", currency.Currency.ChainID, currency.Currency.Address, err)
		}

		coin := api.Coin{
			CurrencyID: id,
			Value:      currency.Value,
		}

		out = append(out, coin)
	}

	return out, nil
}

// createCoinSnapshotList takes a list of coins and transforms the to the API data format,
// translating chain ID and currency address pairs to the Currency ID.
func (a *API) createCoinSnapshotList(snapshots []datapoint.CoinSnapshot) ([]api.CoinSnapshot, error) {

	out := make([]api.CoinSnapshot, 0, len(snapshots))
	for _, snapshot := range snapshots {
		snapshot := snapshot

		coins, err := a.createCoinList(snapshot.Coins)
		if err != nil {
			return nil, fmt.Errorf("could not create coin list: %w", err)
		}

		s := api.CoinSnapshot{
			Value: coins,
			Time:  &snapshot.Date,
		}

		out = append(out, s)
	}

	return out, nil
}

// createValueHistoryRecord creates the API response type for historic data for a stat.
func (a *API) createValueHistoryRecord(id string, snapshots []datapoint.CoinSnapshot) (api.ValueHistory, error) {

	cs, err := a.createCoinSnapshotList(snapshots)
	if err != nil {
		return api.ValueHistory{}, fmt.Errorf("could not create coin snapshot list: %w", err)
	}

	out := api.ValueHistory{
		ID:        id,
		Snapshots: cs,
	}

	return out, nil
}
