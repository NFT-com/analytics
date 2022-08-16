package stats

import (
	"fmt"
	"time"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// TODO: Without comparing NFT prices expressed in a single currency, this functionality makes little sense.
// Currently it only finds the lowest `currency_value` number, which might not be the lowest price.
// See https://github.com/NFT-com/analytics/issues/85

// CollectionLowestPriceHistory returns the lowest price for the collection in the given interval.
func (s *Stats) CollectionLowestPriceHistory(address identifier.Address, from time.Time, to time.Time) ([]datapoint.LowestPrice, error) {

	intervalQuery := s.db.
		Table("sales, LATERAL generate_series(?::timestamp, ?::timestamp, INTERVAL '1 day') AS start_date",
			from.Format(timeFormat),
			to.Format(timeFormat)).
		Select([]string{
			"sales.*",
			"start_date",
			"start_date + interval '1 day' AS end_date"}).
		Where("chain_id = ?", address.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", address.Address)

	seriesQuery := s.db.
		Table("(?) s", intervalQuery).
		Select("s.*").
		Where("s.emitted_at > s.start_date").
		Where("s.emitted_at <= s.end_date")

	query := s.db.
		Table("(?) st", seriesQuery).
		Select("MIN(st.currency_value) AS currency_value, chain_id, LOWER(currency_address) AS currency_address, st.start_date, st.end_date").
		Group("start_date").
		Group("end_date").
		Group("chain_id").
		Group("LOWER(currency_address)").
		Order("start_date DESC")

	var prices []lowestPriceResult
	err := query.Find(&prices).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve lowest price: %w", err)
	}

	out := make([]datapoint.LowestPrice, 0, len(prices))
	for _, p := range prices {

		currency := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: p.ChainID,
				Address: p.Address,
			},
			Amount: p.Amount,
		}

		price := datapoint.LowestPrice{
			Currency: currency,
			Start:    p.Start,
			End:      p.End,
		}

		out = append(out, price)
	}

	return out, nil
}
