package api

import (
	"fmt"

	"github.com/NFT-com/graph-api/graph/models/api"
)

type traitStats struct {
	// ocurrences keeps track of how many times a trait type-value combination is found.
	occurrences map[string]uint

	// knownTraits keeps track of all known traits for a collection, counting how many NFTs have them.
	knownTraits map[string]uint
}

func (ct collectionTraits) stats() traitStats {

	stats := traitStats{
		occurrences: make(map[string]uint),
		knownTraits: make(map[string]uint),
	}

	// Go through all traits found and count number of ocurrence of
	// each individual trait type/value combo, as well as how many times
	// was each trait type seen.
	for _, traitList := range ct {

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
			stats.knownTraits[t]++
		}
	}

	return stats
}

// calcTraitCollectionRarity takes the collection size, the stats about the traits in that collection,
// and a list of traits (key-value pairs) and calculates the individual trait rarity information. It returns
// the overall rarity for such a trait combination, as well as the individual trait rarity info, including
// missing traits.
func calcTraitCollectionRarity(size uint, stats traitStats, traits []*api.Trait) (float64, []*api.Trait) {

	// Traits populated with the rarity score.
	// This list includes traits that are not in the original list of traits,
	// but are known within the collection.
	out := make([]*api.Trait, 0, len(traits))

	// Keep track of all traits we found.
	found := make(map[string]struct{})

	// Overall rarity for this trait combination.
	overallRarity := 1.0

	// For each trait, see how many times that trait with that value occurred in the collection.
	// Divide the number of occurrences with the collection size to see how frequent that trait is.
	for _, trait := range traits {

		found[trait.Type] = struct{}{}

		key := formatTraitKey(trait)
		occurrences := stats.occurrences[key]

		traitRarity := float64(occurrences) / float64(size)

		t := api.Trait{
			Type:   trait.Type,
			Value:  trait.Value,
			Rarity: traitRarity,
		}

		out = append(out, &t)
		// Update the overall rarity.
		overallRarity = overallRarity * traitRarity
	}

	// Go through known traits for this collection. For all traits that are missing,
	// calculate the probability for that.
	for trait, count := range stats.knownTraits {

		// Is this known trait found in this list?
		_, have := found[trait]
		if have {
			continue
		}

		// What is the probability that an NFT does not have this trait?
		// If there's a 100 NFTs in a collection and only one has this trait,
		// the probability is (100 - 1)/100 = 99%
		rarity := float64(size-count) / float64(size)

		missing := api.Trait{
			Type:   trait,
			Value:  "",
			Rarity: rarity,
		}

		out = append(out, &missing)
		// Update the overall rarity.
		overallRarity = overallRarity * rarity
	}

	return overallRarity, out
}

func (ts traitStats) Print() {
	fmt.Printf("occurrences\n")
	for trait, count := range ts.occurrences {
		fmt.Printf("\t%v - %d\n", trait, count)
	}

	fmt.Printf("present traits\n")
	for trait, count := range ts.knownTraits {
		fmt.Printf("\t%v - %d\n", trait, count)
	}
}

func formatTraitKey(trait *api.Trait) string {
	return trait.Type + ":" + trait.Value
}
