package api

import (
	"github.com/NFT-com/graph-api/graph/models/api"
)

func (s *Server) nftTraits(nft *api.NFT, wantRarity bool) ([]*api.Trait, error) {

	// Get trait information for the current NFT.
	traits, err := s.storage.NFTTraits(nft.ID, wantRarity)
	if err != nil {
		s.logError(err).
			Str("nft", nft.ID).
			Msg("could not retrieve NFT traits")
		return nil, errRetrieveTraitsFailed
	}

	// If we don't need to know the NFT rarity, we're done.
	if !wantRarity {
		return traits, nil
	}

	// Find rarity for missing traits.

	foundTraits := make([]string, 0, len(traits))
	for _, trait := range traits {
		foundTraits = append(foundTraits, trait.Type)
	}

	missingTraits, err := s.storage.NFTMissingTraitRatio(nft.Collection, foundTraits)
	if err != nil {
		s.logError(err).
			Str("nft", nft.ID).
			Msg("could not retrieve NFT traits")
		return nil, errRetrieveTraitsFailed
	}

	traits = append(traits, missingTraits...)

	return traits, nil
}

// FIXME: Create a type instead of using map[string][]*api.Trait
func (s *Server) getTraitsForCollection(collectionID string) (map[string][]*api.Trait, error) {

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
