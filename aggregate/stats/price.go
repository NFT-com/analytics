package stats

import (
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
func (s *Stats) NFTBatchPrice(nfts []identifier.Address) (map[identifier.NFT]datapoint.Price, error) {
	return nil, fmt.Errorf("TBD: Not implemented")
}
