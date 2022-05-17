package storage

import (
	"fmt"

	"github.com/NFT-com/graph-api/aggregate/models/identifier"
)

func (s *Storage) NFT(id string) (identifier.NFT, error) {

	var nft nftIdentifier

	err := s.db.
		Table("nfts n").
		Table("collections c").
		Table("networks").
		Where("n.id = ?", id).
		Where("c.id = n.collection_id").
		Where("networks.id = c.network_id").
		First(&nft).Error
	if err != nil {
		return identifier.NFT{}, fmt.Errorf("could not retrieve NFT address: %w", err)
	}

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
