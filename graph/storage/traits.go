package storage

import (
	"fmt"

	"github.com/NFT-com/analytics/graph/models/api"
)

// NFTTraits retrieves a list of traits of that NFT.
func (s *Storage) NFTTraits(nftID string) ([]*api.Trait, error) {

	query := api.Trait{
		NFT: nftID,
	}

	var traits []*api.Trait
	err := s.db.Where(query).Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve traits: %w", err)
	}

	return traits, nil
}

// CollectionTraits retrieves all of the traits belonging to NFTs in a specific collection.
func (s *Storage) CollectionTraits(collectionID string) ([]*api.Trait, error) {

	var traits []*api.Trait
	err := s.db.
		Table("traits_collections").
		Where("collection_id = ?", collectionID).
		Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve traits for collections: %w", err)
	}

	return traits, nil
}
