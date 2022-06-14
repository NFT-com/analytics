package api

import (
	"context"

	"github.com/NFT-com/analytics/graph/models/api"
	"github.com/NFT-com/analytics/graph/stats/collection"
)

// getNFT returns a single NFT based on its ID.
func (s *Server) getNFT(ctx context.Context, id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return s.getNFTDetails(ctx, nft)
}

// getNFTByTokenID returns a single NFT based on the combination of networkID, contract address and token ID.
func (s *Server) getNFTByTokenID(ctx context.Context, networkID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(networkID, contract, tokenID)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Str("contract", contract).
			Str("token_id", tokenID).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return s.getNFTDetails(ctx, nft)
}

// getNFTDetails retrieves the NFT rarity and/or trait information.
func (s *Server) getNFTDetails(ctx context.Context, nft *api.NFT) (*api.NFT, error) {

	// Parse the query to know how much information to return/calculate.
	req := parseNFTQuery(ctx)

	// If we do not need traits nor rarity, we're done.
	if !req.traits && !req.needRarity() {
		return nft, nil
	}

	// If we do not need rarity, just fetch the traits for this NFT.
	if !req.needRarity() {
		traits, err := s.storage.NFTTraits(nft.ID)
		if err != nil {
			s.logError(err).Str("id", nft.ID).Msg("could not retrieve traits")
			return nil, errRetrieveTraitsFailed
		}

		nft.Traits = traits
		return nft, nil
	}

	// Get traits and calculate rarity.
	traits, err := s.getTraitsForCollection(nft.Collection)
	if err != nil {
		return nil, errRetrieveTraitsFailed
	}
	nft.Traits = traits[nft.ID]

	// Get the size of this collection.
	size, err := s.storage.CollectionSize(nft.Collection)
	if err != nil {
		s.logError(err).Str("collection", nft.Collection).Msg("could not retrieve collection size")
		return nil, errRetrieveTraitsFailed
	}

	stats := traits.CalculateStats()
	rarity, traitRarity := stats.CalculateRarity(size, nft.Traits)

	nft.Rarity = rarity

	// Returned traits include traits missing for this NFT.
	// Only set them if individual trait rarity is requested.
	if req.traitRarity {
		nft.Traits = traitRarity
	}

	return nft, nil
}

// nfts returns a list of NFTs fitting the search criteria.
func (s *Server) nfts(ctx context.Context, owner *string, collectionID *string, rarityMax *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collectionID, orderBy, s.searchLimit)
	if err != nil {
		log := s.logError(err)
		if owner != nil {
			log = log.Str("owner", *owner)
		}
		if collectionID != nil {
			log = log.Str("collection", *collectionID)
		}
		if rarityMax != nil {
			log = log.Float64("max_rarity", *rarityMax)
		}
		log.Msg("could not retrieve nfts")
		return nil, errRetrieveNFTFailed
	}

	// Parse the query to know how much information to return/calculate.
	req := parseNFTQuery(ctx)

	// If we do not need traits nor rarity, we're done.
	if !req.traits && !req.needRarity() {
		return nfts, nil
	}

	// We only need traits and no rarity stats.
	if !req.needRarity() {

		// Map collection ID to collection traits.
		traits := make(map[string]collection.TraitMap)

		for _, nft := range nfts {
			// Lookup traits for this collection.
			// If we don't have them already cached, fetch them now.
			ctraits, ok := traits[nft.Collection]
			if !ok {
				tc, err := s.getTraitsForCollection(nft.Collection)
				if err != nil {
					s.logError(err).Str("collection", nft.Collection).Msg("could not retrieve traits for collection")
					return nil, errRetrieveTraitsFailed
				}
				traits[nft.Collection] = tc
				ctraits = tc
			}

			nft.Traits = ctraits[nft.ID]
		}

		return nfts, nil
	}

	// We need rarity calculations.

	// Cache traits and stats for each of the collections.
	traits := make(map[string]collection.TraitMap)
	stats := make(map[string]collection.Stats)
	sizes := make(map[string]uint)

	for _, nft := range nfts {
		// Lookup stats for this collection.
		// If we don't have them already cached, fetch them now, calculate stats and cache them.
		cstats, ok := stats[nft.Collection]
		if !ok {
			tc, err := s.getTraitsForCollection(nft.Collection)
			if err != nil {
				s.logError(err).Str("collection", nft.Collection).Msg("could not retrieve traits for collection")
				return nil, errRetrieveTraitsFailed
			}

			size, err := s.storage.CollectionSize(nft.Collection)
			if err != nil {
				s.logError(err).Str("collection", nft.Collection).Msg("could not retrieve collection size")
				return nil, errRetrieveTraitsFailed
			}

			traits[nft.Collection] = tc
			st := tc.CalculateStats()
			stats[nft.Collection] = st
			sizes[nft.Collection] = size

			cstats = st
		}

		nft.Traits = traits[nft.Collection][nft.ID]
		rarity, traitRarity := cstats.CalculateRarity(sizes[nft.Collection], nft.Traits)
		nft.Rarity = rarity
		if req.traitRarity {
			nft.Traits = traitRarity
		}
	}

	return nfts, nil
}
