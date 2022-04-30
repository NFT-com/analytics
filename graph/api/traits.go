package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

// collectionTraits represents a map of NFT IDs to NFT traits.
type collectionTraits map[string][]*api.Trait

// FIXME: Create a type instead of using map[string][]*api.Trait
func (s *Server) getTraitsForCollection(collectionID string) (collectionTraits, error) {

	// Get traits for collection.
	list, err := s.storage.CollectionTraits(collectionID)
	if err != nil {
		s.logError(err).
			Str("collection", collectionID).
			Msg("could not retrieve collection traits")
		return nil, errRetrieveTraitsFailed
	}

	// Transform the list of traits to a map, mapping NFT ID to a list of traits.
	traits := make(map[string][]*api.Trait)
	for _, trait := range list {
		trait := trait

		t, ok := traits[trait.NFT]
		if ok {
			t = append(t, trait)
			traits[trait.NFT] = t
			continue
		}

		nftTraits := make([]*api.Trait, 0)
		nftTraits = append(nftTraits, trait)
		traits[trait.NFT] = nftTraits
	}

	return traits, nil
}
