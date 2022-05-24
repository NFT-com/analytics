package stats

import (
	"errors"
	"fmt"

	"github.com/NFT-com/graph-api/aggregate/models/datapoint"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

// NFTPrice returns the current NFT price for an NFT.
func (s *Stats) NFTPrice(nft identifier.NFT) (datapoint.Price, error) {

	query := s.db.
		Table("sales").
		Select("trade_price").
		Where("chain_id = ?", nft.Collection.ChainID).
		Where("collection_address = ?", nft.Collection.Address).
		Where("token_id = ?", nft.TokenID).
		Order("emitted_at DESC").
		Limit(1)

	var price datapoint.Price
	err := query.Take(&price).Error
	if err != nil {
		return datapoint.Price{}, fmt.Errorf("could not retrieve price: %w", err)
	}

	return price, nil
}

// NFTBatchPrice returns the list of prices for the specified NFTs.
func (s *Stats) NFTBatchPrices(nfts []identifier.NFT) (map[identifier.NFT]datapoint.Price, error) {

	if len(nfts) == 0 {
		return nil, errors.New("id list must be non-empty")
	}

	selectFields := []string{
		"chain_id",
		"collection_address",
		"token_id",
		"trade_price",
		"row_number() OVER (PARTITION BY chain_id, collection_address, token_id ORDER BY emitted_at DESC) AS rank",
	}

	priceQuery := s.db.
		Table("sales").
		Select(selectFields)

	filter := s.createNFTFilter(nfts)
	priceQuery = priceQuery.Where(filter)

	// filterQuery selects only the latest prices from the priceQuery result.
	filterQuery := s.db.
		Table("( ? ) p", priceQuery).
		Where("rank = 1")

	// Get the list of prices.
	var prices []batchPriceResult
	err := filterQuery.Find(&prices).Error
	if err != nil {
		return nil, fmt.Errorf("could not get prices: %w", err)
	}

	// Transform the list of prices into a map, mapping the NFT identifier to the price point.
	priceMap := make(map[identifier.NFT]datapoint.Price, len(nfts))
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

		// Price record.
		nftPrice := datapoint.Price{
			Price: price.TradePrice,
		}

		priceMap[nft] = nftPrice
	}

	return priceMap, nil
}
