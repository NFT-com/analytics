package collection

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

// Stats contains trait statistics for a given collection.
type Stats struct {
	// Ocurrences keeps track of how many times a trait type-value combination is found.
	Occurrences map[string]uint

	// KnownTraits keeps track of all known traits for a collection, counting how many NFTs have them.
	KnownTraits map[string]uint
}

// CalculateStats calculates trait frequency for the trait map, as well as keeping a list of all known traits
// within that collection.
func (t TraitMap) CalculateStats() Stats {

	s := Stats{
		Occurrences: make(map[string]uint),
		KnownTraits: make(map[string]uint),
	}

	// Go through all traits found and count number of occurrences of
	// each individual trait type/value combo, as well as how many times
	// was each trait type seen.
	for _, traits := range t {

		// We want to keep track of how many NFTs have a certain trait.
		// However, if there's an NFT that has the same trait twice,
		// we don't want to count that.
		distinct := make(map[string]struct{})
		for _, trait := range traits {
			key := formatTraitKey(trait)
			s.Occurrences[key]++

			distinct[trait.Name] = struct{}{}
		}

		// Add all distinct trait types to the trait type counter.
		for t := range distinct {
			s.KnownTraits[t]++
		}
	}

	return s
}

func formatTraitKey(trait *api.Trait) string {
	return trait.Name + ":" + trait.Value
}
