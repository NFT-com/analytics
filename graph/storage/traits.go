package storage

import (
	"fmt"

	"github.com/NFT-com/graph-api/graph/models/api"
)

const (
	traitRatioFunctionName = `traits_ratio`
)

// NFTTraits retrieves a list of traits that the specified NFT has,
// along with the ratio describing how many NFTs from that collection
// have the specific trait name/value combination.
func (s *Storage) NFTTraits(id string, calculateRarity bool) ([]*api.Trait, error) {

	if calculateRarity {
		return s.nftTraitRarity(id)
	}

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

func (s *Storage) nftTraitRarity(id string) ([]*api.Trait, error) {

	// Prepare the SQL statement.
	query := fmt.Sprintf("SELECT * FROM %s(?)", traitRatioFunctionName)

	// Execute the SQL query.
	var traits []*api.Trait
	err := s.db.Raw(query, id).Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve trait data: %w", err)
	}

	return traits, nil
}

// NFTMissingTraitRatio determines what is the probability that an NFT does NOT have a specific trait.
// It accepts a list of traits that an NFT has, and sees what other traits can be found in a certain
// collection.
func (s *Storage) NFTMissingTraitRatio(collectionID string, foundTraits []string) ([]*api.Trait, error) {

	db := s.db.
		Table("traits_collections tc").
		Select("name, 1 - COUNT(DISTINCT(nft))::NUMERIC / (SELECT COUNT(*)::NUMERIC FROM nfts WHERE collection = ?) as ratio",
			collectionID,
		).Where("name not in (?)", foundTraits).
		Where("tc.collection = ?", collectionID).
		Group("name")

	var traits []*api.Trait
	err := db.Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not calculate missing trait rarity: %w", err)
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
