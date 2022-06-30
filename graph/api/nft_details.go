package api

import (
	"fmt"

	"github.com/NFT-com/analytics/graph/api/internal/query"
	"github.com/NFT-com/analytics/graph/models/api"
)

func (s *Server) expandNFTDetails(query *query.NFT, nft *api.NFT) error {

	// Retrieve owner if it was requested.
	if query.Owners {
		owners, err := s.storage.NFTOwners(nft.ID)
		if err != nil {
			return fmt.Errorf("could not retrieve owners for the NFT: %w", err)
		}
		nft.Owners = owners
	}

	// Get NFT stats from the aggregation API.
	err := s.getNFTStats(query, nft)
	if err != nil {
		// Continue even if stats could not be retrieved (e.g. API was unavailable).
		s.log.Error().
			Err(err).
			Str("id", nft.ID).
			Msg("could not retrieve NFT stats")
	}

	// Get trait data and, if required, calculate rarity as well.
	err = s.getNFTRarityAndTraitData(query, nft)
	if err != nil {
		return fmt.Errorf("could not retrieve rarity/trait information: %w", err)
	}

	return nil
}

// getNFTRarityAndTraitData retrieves the NFT rarity and/or trait information.
func (s *Server) getNFTRarityAndTraitData(query *query.NFT, nft *api.NFT) error {

	// If we do not need traits nor rarity, we're done.
	if !query.Traits && !query.NeedRarity() {
		return nil
	}

	// If we do not need rarity, just fetch the traits for this NFT.
	if !query.NeedRarity() {
		traits, err := s.storage.NFTTraits(nft.ID)
		if err != nil {
			s.logError(err).Str("id", nft.ID).Msg("could not retrieve traits")
			return errRetrieveTraitsFailed
		}

		nft.Traits = traits
		return nil
	}

	// Get traits and calculate rarity.
	traits, err := s.getTraitsForCollection(nft.Collection)
	if err != nil {
		return errRetrieveTraitsFailed
	}
	nft.Traits = traits[nft.ID]

	// Get the size of this collection.
	size, err := s.storage.CollectionSize(nft.Collection)
	if err != nil {
		s.logError(err).Str("collection", nft.Collection).Msg("could not retrieve collection size")
		return errRetrieveTraitsFailed
	}

	stats := traits.CalculateStats()
	rarity, traitRarity := stats.CalculateRarity(size, nft.Traits)

	nft.Rarity = rarity

	// Returned traits include traits missing for this NFT.
	// Only set them if individual trait rarity is requested.
	if query.TraitRarity {
		nft.Traits = traitRarity
	}

	return nil
}
