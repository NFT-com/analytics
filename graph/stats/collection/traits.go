package collection

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

// TraitMap represents a map of collection traits, mapping NFT IDs to NFT traits.
type TraitMap map[string][]*api.Trait

// NewTraitMap creates a new collection trait map from the list of NFT traits.
func NewTraitMap(list []*api.Trait) TraitMap {

	traits := make(map[string][]*api.Trait)

	for _, trait := range list {
		trait := trait

		t, ok := traits[trait.NFT]
		if ok {
			t = append(t, trait)
			traits[trait.NFT] = t
			continue
		}

		traits[trait.NFT] = []*api.Trait{trait}
	}

	return traits
}
