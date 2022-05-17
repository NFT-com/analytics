package storage

import (
	"fmt"

	aggregate "github.com/NFT-com/graph-api/aggregate/api"
	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

func (s *Storage) NFT(id string) (identifier.NFT, error) {

	// Note: Using `Find` with a limit of 1 instead of `First` because the generated SQL
	// uses the wrong table name otherwise.

	var nfts []nftIdentifier
	err := s.db.
		Table("nfts n, collections c, networks").
		Select("networks.chain_id, c.contract_address, n.token_id").
		Where("n.id = ?", id).
		Where("c.id = n.collection_id").
		Where("networks.id = c.network_id").
		Limit(1).
		Find(&nfts).Error
	if err != nil {
		return identifier.NFT{}, fmt.Errorf("could not retrieve NFT address: %w", err)
	}
	if len(nfts) == 0 {
		return identifier.NFT{}, aggregate.ErrRecordNotFound
	}

	nft := nfts[0]

	collection := identifier.Address{
		ChainID: nft.ChainID,
		Address: nft.Address,
	}

	nftAddress := identifier.NFT{
		Collection: collection,
		TokenID:    nft.TokenID,
	}

	return nftAddress, nil
}
