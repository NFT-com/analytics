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
