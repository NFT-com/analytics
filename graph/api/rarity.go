package api

import (
	"fmt"

	"github.com/NFT-com/graph-api/graph/models/api"
)

type traitStats struct {
	// ocurrences keeps track of how many times a trait type-value combination is found.
	occurrences map[string]uint

	// presentTraits keeps track of how many NFTs have a specific trait type at all.
	presentTraits map[string]uint
}

func extractTraitStats(traits map[string][]*api.Trait) traitStats {

	stats := traitStats{
		occurrences:   make(map[string]uint),
		presentTraits: make(map[string]uint),
	}

	// Go through all traits found and count number of ocurrence of
	// each individual trait type/value combo, as well as how many times
	// was each trait type seen.
	for _, traitList := range traits {

		// We want to keep track of how many NFTs have a certain trait.
		// However, if there's an NFT that has the same trait twice,
		// we don't want to count that.
		distinctTraits := make(map[string]struct{})

		for _, t := range traitList {
			key := formatTraitKey(t)
			stats.occurrences[key]++

			distinctTraits[t.Type] = struct{}{}
		}

		// Add all distinct trait types to the trait type counter.
		for t := range distinctTraits {
			stats.presentTraits[t]++
		}
	}

	return stats
}

func (ts traitStats) Print() {
	fmt.Printf("occurrences\n")
	for trait, count := range ts.occurrences {
		fmt.Printf("\t%v - %d\n", trait, count)
	}

	fmt.Printf("present traits\n")
	for trait, count := range ts.presentTraits {
		fmt.Printf("\t%v - %d\n", trait, count)
	}
}

func formatTraitKey(trait *api.Trait) string {
	return trait.Type + ":" + trait.Value
}
