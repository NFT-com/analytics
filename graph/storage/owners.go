package storage

import (
	"fmt"

	"github.com/NFT-com/analytics/aggregate/models/identifier"
	"github.com/NFT-com/analytics/graph/models/api"
)

// NFTOwners returns the addresses of all accounts owning this NFT.
func (s *Storage) NFTOwners(nftID string) ([]api.Owner, error) {

	var owners []api.Owner
	err := s.db.
		Table("owners").
		Select("owner, SUM(number) AS number").
		Where("owner != ?", identifier.ZeroAddress).
		Where("nft_id = ?", nftID).
		Group("owner").
		Having("SUM(number) > ?", 0).
		Find(&owners).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT owners: %w", err)
	}

	return owners, nil
}

// CollectionOwners returns the list of owners for each NFT in the collection.
func (s *Storage) CollectionOwners(collectionID string) (map[string][]api.Owner, error) {

	var owners []api.Owner
	err := s.db.
		Table("owners o, nfts n").
		Select("o.owner, o.nft_id, SUM(o.number) as number").
		Where("o.owner != ?", identifier.ZeroAddress).
		Where("o.nft_id = n.id").
		Where("n.collection_id = ?", collectionID).
		Where("n.deleted != TRUE").
		Group("owner, nft_id").
		Having("SUM(o.number) > ?", 0).
		Find(&owners).Error
	if err != nil {
		return nil, fmt.Errorf("could not lookup owners for a collection: %w", err)
	}

	// Map owners to NFT IDs.
	out := make(map[string][]api.Owner)
	for _, owner := range owners {
		list, ok := out[owner.NFTID]
		if ok {
			list = append(list, owner)
			out[owner.NFTID] = list
			continue
		}

		list = make([]api.Owner, 0, 1)
		list = append(list, owner)

		out[owner.NFTID] = list
	}

	return out, nil
}

// nftListOwners retrieves owners for a list of NFTs.
func (s *Storage) nftListOwners(nftIDs []string) (map[string][]api.Owner, error) {

	var owners []api.Owner
	err := s.db.
		Table("owners o, nfts n").
		Select("o.owner, o.nft_id, SUM(o.number) AS number").
		Where("o.owner != ?", identifier.ZeroAddress).
		Where("o.nft_id = n.id").
		Where("n.id IN (?)", nftIDs).
		Group("owner, nft_id").
		Having("SUM(o.number) > ?", 0).
		Find(&owners).Error
	if err != nil {
		return nil, fmt.Errorf("could not lookup owners for a collection: %w", err)
	}

	// Map owners to NFT IDs.
	out := make(map[string][]api.Owner)
	for _, owner := range owners {
		list, ok := out[owner.NFTID]
		if ok {
			list = append(list, owner)
			out[owner.NFTID] = list
			continue
		}

		list = make([]api.Owner, 0, 1)
		list = append(list, owner)

		out[owner.NFTID] = list
	}

	return out, nil
}
