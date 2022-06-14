package collection

import (
	"github.com/NFT-com/analytics/graph/models/api"
)

// CalculateRarity takes the collection size and a list of traits (key-value pairs)
// and calculates the individual trait rarity information. It returns the overall rarity
// for such a trait combination, as well as the individual trait rarity info, which includes
// missing traits.
func (s *Stats) CalculateRarity(total uint, traits []*api.Trait) (float64, []*api.Trait) {

	// Traits populated with the rarity score.
	// This list includes traits that are not in the original list of traits,
	// but are known within this trait map (typically a collection).
	out := make([]*api.Trait, 0, len(traits))

	// Keep track of all traits we found.
	found := make(map[string]struct{})

	// Overall rarity for this trait combination.
	overallRarity := 1.0

	// For each trait, see how many times that trait with that value occurred in the collection.
	// Divide the number of occurrences with the collection size to see how frequent that trait is.
	for _, trait := range traits {

		found[trait.Name] = struct{}{}

		key := formatTraitKey(trait)
		occurrences := s.Occurrences[key]

		rarity := float64(occurrences) / float64(total)

		t := api.Trait{
			Name:   trait.Name,
			Value:  trait.Value,
			Rarity: rarity,
		}

		out = append(out, &t)

		// Update the overall rarity.
		overallRarity = overallRarity * rarity
	}

	// Go through known traits for this collection. For all traits that are missing,
	// calculate the probability for that.
	for trait, count := range s.KnownTraits {

		// Is this known trait found in this list?
		_, have := found[trait]
		if have {
			continue
		}

		// What is the probability that an NFT does not have this trait?
		// If there's a 100 NFTs in a collection and only one has this trait,
		// the probability is (100 - 1)/100 = 99%
		rarity := float64(total-count) / float64(total)

		missing := api.Trait{
			Name:   trait,
			Value:  "",
			Rarity: rarity,
		}

		out = append(out, &missing)

		// Update the overall rarity.
		overallRarity = overallRarity * rarity
	}

	return overallRarity, out
}
