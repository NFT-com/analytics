package storage

import (
	"fmt"

	"github.com/NFT-com/graph-api/graph/models/api"
)

const (
	traitRatioFunctionName = `traits_ratio`
)

// traitRatio represents the trait ratio function query result.
type traitRatio struct {
	Name  string  `gorm:"column:name"`
	Value string  `gorm:"column:value"`
	Ratio float64 `gorm:"column:ratio"`
}

// NFTTraitRatio retrieves a list of traits that the specified NFT has,
// along with the ratio describing how many NFTs from that collection
// have the specific trait name/value combination.
func (s *Storage) NFTTraitRatio(id string) ([]*api.TraitRatio, error) {

	// Prepare the SQL statement.
	query := fmt.Sprintf("SELECT * FROM %s(?)", traitRatioFunctionName)

	// Execute the SQL query.
	var traits []traitRatio
	err := s.db.Raw(query, id).Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve trait data: %w", err)
	}

	// Translate the query result to the expected format.
	out := make([]*api.TraitRatio, 0, len(traits))
	for _, t := range traits {

		trait := api.TraitRatio{
			Trait: api.Trait{
				Type:  t.Name,
				Value: t.Value,
			},
			Ratio: t.Ratio,
		}

		out = append(out, &trait)
	}

	return out, nil
}

// NFTMissingTraitRatio determines what is the probability that an NFT does NOT have a specific trait.
// It will accept a list of traits that an NFT has, and see what other traits can be found in a certain
// collection.
func (s *Storage) NFTMissingTraitRatio(collectionID string, foundTraits []string) ([]*api.TraitRatio, error) {

	// FIXME: The select might be a subquery
	db := s.db.
		Table("traits_collections tc").
		Select("name, 1 - COUNT(DISTINCT(nft))::NUMERIC / (SELECT COUNT(*)::NUMERIC FROM nfts WHERE collection = ?) as ratio",
			collectionID,
		).Where("name not in (?)", foundTraits).
		Where("tc.collection = ?", collectionID).
		Group("name")

	var traits []traitRatio
	err := db.Find(&traits).Error
	if err != nil {
		return nil, fmt.Errorf("could not calculate missing trait rarity: %w", err)
	}

	// Translate the query result to the expected format.
	out := make([]*api.TraitRatio, 0, len(traits))
	for _, t := range traits {

		trait := api.TraitRatio{
			Trait: api.Trait{
				Type:  t.Name,
				Value: t.Value,
			},
			Ratio: t.Ratio,
		}

		out = append(out, &trait)
	}

	return out, nil
}
