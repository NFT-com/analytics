package stats

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/datapoint"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// NFTPrice returns the current NFT price for an NFT.
func (s *Stats) NFTPrice(nft identifier.NFT) ([]datapoint.Coin, error) {

	// NOTE: If we start supporting scenarios with multiple payments in a single sale (e.g. 10 ETH + 1 WETH for an NFT)
	// this approach will need to be altered.

	query := s.db.
		Table("sales").
		Select("currency_value, chain_id, LOWER(currency_address) AS currency_address").
		Where("chain_id = ?", nft.Collection.ChainID).
		Where("LOWER(collection_address) = LOWER(?)", nft.Collection.Address).
		Where("token_id = ?", nft.TokenID).
		Order("emitted_at DESC").
		Limit(1)

	var res priceResult
	err := query.Find(&res).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve price: %w", err)
	}

	price := []datapoint.Coin{
		{
			Currency: identifier.Currency{
				ChainID: res.ChainID,
				Address: res.Address,
			},
			Value: res.Value,
		},
	}

	return price, nil
}

// CollectionPrices returns the list of prices for NFTs in a specified collection.
// Prices are mapped to the NFT identifier, with the collection contract address being lowercased.
// NOTE: CollectionPrices and CollectionAveragePrices could return a map where the keys are simple strings - token IDs,
// since all of the tokens are from the same collection. However, for uniformity with the rest of the package, they use `identifier.NFT` for mapping.
func (s *Stats) CollectionPrices(address identifier.Address) (map[identifier.NFT][]datapoint.Coin, error) {

	selectFields := []string{
		"chain_id",
		"LOWER(collection_address) AS collection_address",
		"token_id",
		"currency_value",
		"LOWER(currency_address) AS currency_address",
		"row_number() OVER (PARTITION BY chain_id, LOWER(collection_address), token_id ORDER BY emitted_at DESC) AS rank",
	}

	filter := s.createCollectionFilter([]identifier.Address{address})

	priceQuery := s.db.
		Table("sales").
		Select(selectFields).
		Where(filter)

	query := s.db.
		Table("( ? ) p", priceQuery).
		Where("rank = 1")

	// Get the list of prices.
	var prices []batchPriceResult
	err := query.Find(&prices).Error
	if err != nil {
		return nil, fmt.Errorf("could not get prices: %w", err)
	}

	// Transform the list of prices into a map, mapping the NFT identifier to the price point.
	priceMap := make(map[identifier.NFT][]datapoint.Coin, len(prices))
	for _, price := range prices {

		// Create the NFT identifier.
		collection := identifier.Address{
			ChainID: price.ChainID,
			Address: price.CollectionAddress,
		}
		nft := identifier.NFT{
			Collection: collection,
			TokenID:    price.TokenID,
		}

		p := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: price.ChainID,
				Address: price.CurrencyAddress,
			},
			Value: price.CurrencyValue,
		}

		_, ok := priceMap[nft]
		if !ok {
			priceMap[nft] = make([]datapoint.Coin, 0)
		}

		priceMap[nft] = append(priceMap[nft], p)
	}

	return priceMap, nil
}

// CollectionAveragePrices returns the list of average prices for NFTs in a specified collection.
// Prices are mapped to the NFT identifier, with the collection contract address being lowercased.
func (s *Stats) CollectionAveragePrices(address identifier.Address) (map[identifier.NFT][]datapoint.Coin, error) {

	selectFields := []string{
		"chain_id",
		"LOWER(collection_address) AS collection_address",
		"token_id",
		"AVG(currency_value) AS currency_value",
		"LOWER(currency_address) AS currency_address",
	}

	filter := s.createCollectionFilter([]identifier.Address{address})

	query := s.db.
		Table("sales").
		Select(selectFields).
		Where(filter).
		Group("chain_id, LOWER(collection_address), token_id, LOWER(currency_address)")

	// Get the list of prices.
	var prices []batchPriceResult
	err := query.Find(&prices).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve prices: %w", err)
	}

	// Transform the list of prices into a map, mapping the NFT identifier to the price point.
	priceMap := make(map[identifier.NFT][]datapoint.Coin, len(prices))
	for _, price := range prices {

		// Create the NFT identifier.
		collection := identifier.Address{
			ChainID: price.ChainID,
			Address: price.CollectionAddress,
		}
		nft := identifier.NFT{
			Collection: collection,
			TokenID:    price.TokenID,
		}

		currency := datapoint.Coin{
			Currency: identifier.Currency{
				ChainID: price.ChainID,
				Address: price.CollectionAddress,
			},
			Value: price.CurrencyValue,
		}

		// If we already have average price for this nft (for some currencies)
		// just append the data for this currency.
		_, ok := priceMap[nft]
		if ok {
			priceMap[nft] = append(priceMap[nft], currency)
			continue
		}

		p := make([]datapoint.Coin, 0)
		p = append(p, currency)
		priceMap[nft] = p
	}

	return priceMap, nil
}
