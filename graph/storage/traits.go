package storage

import (
	"fmt"

	"github.com/NFT-com/graph-api/graph/models/api"
)

// NFTTraits retrieves a list of traits of that NFT.
func (s *Storage) NFTTraits(id string) ([]*api.Trait, error) {

	var traits []*api.Trait
	err := s.db.
		Select("*").
		Table("traits").
		Where("nft = ?", id).
		Find(&traits).Error
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
		Where("collection = ?", collectionID).
		Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT traits: %w", err)
	}

	return traits, nil
}
