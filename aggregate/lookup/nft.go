package lookup

import (
	"fmt"

	aggregate "github.com/NFT-com/analytics/aggregate/api"
	"github.com/NFT-com/analytics/aggregate/models/identifier"
)

// NFT returns the identifier of the specified NFT.
func (l *Lookup) NFT(id string) (identifier.NFT, error) {

	// NOTE: Using `Find` with a limit of 1 instead of `First` because the generated SQL
	// uses the wrong table name otherwise.

	query := l.db.
		Table("nfts n, collections c, networks").
		Select("networks.chain_id, c.contract_address, n.token_id").
		Where("n.id = ?", id).
		Where("c.id = n.collection_id").
		Where("networks.id = c.network_id").
		Limit(1)

	var nfts []nftIdentifier
	err := query.Find(&nfts).Error
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

// NFTs returns a set of NFT identifiers, mapped by their IDs.
func (l *Lookup) NFTs(ids []string) (map[string]identifier.NFT, error) {

	query := l.db.
		Table("nfts n, collections c, networks").
		Select("n.id, networks.chain_id, c.contract_address, n.token_id").
		Where("c.id = n.collection_id").
		Where("networks.id = c.network_id").
		Where("n.id IN ?", ids)

	var nfts []nftIdentifier
	err := query.Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT addresses: %w", err)
	}

	addresses := make(map[string]identifier.NFT, len(nfts))

	for _, nft := range nfts {

		collection := identifier.Address{
			ChainID: nft.ChainID,
			Address: nft.Address,
		}

		nftAddress := identifier.NFT{
			Collection: collection,
			TokenID:    nft.TokenID,
		}

		addresses[nft.ID] = nftAddress
	}

	return addresses, nil
}
