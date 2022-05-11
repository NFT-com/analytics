package api

import (
	"github.com/NFT-com/graph-api/graph/stats/collection"
)

func (s *Server) getTraitsForCollection(collectionID string) (collection.TraitMap, error) {

	// Get traits for collection.
	list, err := s.storage.CollectionTraits(collectionID)
	if err != nil {
		s.logError(err).
			Str("collection", collectionID).
			Msg("could not retrieve collection traits")
		return nil, errRetrieveTraitsFailed
	}

	// Create a trait map for this collection.
	traits := collection.CreateTraitMap(list)

	return traits, nil
}
