package storage

import (
	"fmt"
)

// NFTOwners returns the addresses of all accounts owning this NFT.
func (s *Storage) NFTOwners(nftID string) ([]string, error) {

	var owners []string
	err := s.db.
		Table("owners").
		Select("owner").
		Where("nft_id = ?", nftID).
		Where("number > 0").
		Find(&owners).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT owners: %w", err)
	}

	return owners, nil
}

// nftOwner models the result of the owner query.
type nftOwner struct {
	Owner string `gorm:"column:owner"`
	NFTID string `gorm:"column:nft_id"`
}

// CollectionOwners returns the list of owners for each NFT in the collection.
func (s *Storage) CollectionOwners(collectionID string) (map[string][]string, error) {

	var owners []nftOwner
	err := s.db.
		Table("owners o, nfts n").
		Select("o.owner, o.nft_id").
		Where("o.nft_id = n.id").
		Where("o.number > 0").
		Where("n.collection_id = ?", collectionID).
		Find(&owners).Error
	if err != nil {
		return nil, fmt.Errorf("could not lookup owners for a collection: %w", err)
	}

	// Map owners to NFT IDs.
	out := make(map[string][]string)
	for _, owner := range owners {
		list, ok := out[owner.NFTID]
		if ok {
			list = append(list, owner.Owner)
			out[owner.NFTID] = list
			continue
		}

		list = make([]string, 0, 1)
		list = append(list, owner.Owner)

		out[owner.NFTID] = list
	}

	return out, nil
}

// nftListOwners retrieves owners for a list of NFTs.
func (s *Storage) nftListOwners(nftIDs []string) (map[string][]string, error) {

	var owners []nftOwner
	err := s.db.
		Table("owners o, nfts n").
		Select("o.owner, o.nft_id").
		Where("o.nft_id = n.id").
		Where("o.number > 0").
		Where("n.id IN (?)", nftIDs).
		Find(&owners).Error
	if err != nil {
		return nil, fmt.Errorf("could not lookup owners for a collection: %w", err)
	}

	// Map owners to NFT IDs.
	out := make(map[string][]string)
	for _, owner := range owners {
		list, ok := out[owner.NFTID]
		if ok {
			list = append(list, owner.Owner)
			out[owner.NFTID] = list
			continue
		}

		list = make([]string, 0, 1)
		list = append(list, owner.Owner)

		out[owner.NFTID] = list
	}

	return out, nil
}
