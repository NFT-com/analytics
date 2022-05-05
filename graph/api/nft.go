package api

import (
	"context"

	"github.com/NFT-com/graph-api/graph/models/api"
	"github.com/NFT-com/graph-api/graph/query"
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

// getNFTByTokenID returns a single NFT based on the combination of chainID, contract address and token ID.
func (s *Server) getNFTByTokenID(ctx context.Context, chainID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(chainID, contract, tokenID)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Str("contract", contract).
			Str("token_id", tokenID).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return s.getNFTDetails(ctx, nft)
}

// getNFTDetails will retrieve the NFT rarity and/or trait information.
func (s *Server) getNFTDetails(ctx context.Context, nft *api.NFT) (*api.NFT, error) {

	// Get the list of selected fields to know how much information to return/calculate.
	sel := query.GetSelection(ctx)

	includeTraits := sel.Has(query.FieldPath(traitField))
	includeTraitRarity := sel.Has(query.FieldPath(traitField, rarityField))
	includeRarity := sel.Has(query.FieldPath(rarityField))

	needRarity := includeRarity || includeTraitRarity

	// If we do not need traits nor rarity, we're done.
	if !includeTraits && !needRarity {
		return nft, nil
	}

	// If we do not need rarity, just fetch the traits for this NFT.
	if !needRarity {
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

	stats := traits.stats()
	rarity, traitRarity := calcTraitCollectionRarity(size, stats, nft.Traits)

	nft.Rarity = rarity

	// Returned traits include traits missing for this NFT.
	// Only set them if individual trait rarity is requested.
	if includeTraitRarity {
		nft.Traits = traitRarity
	}

	return nft, nil
}

// nfts returns a list of NFTs fitting the search criteria.
func (s *Server) nfts(ctx context.Context, owner *string, collection *string, rarityMax *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collection, orderBy, s.searchLimit)
	if err != nil {
		log := s.logError(err)
		if owner != nil {
			log = log.Str("owner", *owner)
		}
		if collection != nil {
			log = log.Str("collection", *collection)
		}
		if rarityMax != nil {
			log = log.Float64("max_rarity", *rarityMax)
		}
		log.Msg("could not retrieve nfts")
		return nil, errRetrieveNFTFailed
	}

	// Get the list of selected fields to know how much information to return/calculate.
	sel := query.GetSelection(ctx)

	includeTraits := sel.Has(query.FieldPath(traitField))
	includeTraitRarity := sel.Has(query.FieldPath(traitField, rarityField))
	includeRarity := sel.Has(query.FieldPath(rarityField))

	needRarity := includeRarity || includeTraitRarity

	// If we do not need traits nor rarity, we're done.
	if !includeTraits && !needRarity {
		return nfts, nil
	}

	// We only need traits and no rarity stats.
	if includeTraits && !needRarity {

		// Map collection ID to collection traits.
		traits := make(map[string]collectionTraits)

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
	traits := make(map[string]collectionTraits)
	stats := make(map[string]traitStats)
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
			st := tc.stats()
			stats[nft.Collection] = st
			sizes[nft.Collection] = size

			cstats = st
		}

		nft.Traits = traits[nft.Collection][nft.ID]
		rarity, traitRarity := calcTraitCollectionRarity(sizes[nft.Collection], cstats, nft.Traits)
		nft.Rarity = rarity
		if includeTraitRarity {
			nft.Traits = traitRarity
		}
	}

	return nfts, nil
}
