package api

import (
	"context"
	"sort"

	"github.com/NFT-com/graph-api/graph/models/api"
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
	query := getQuerySelection(ctx)

	includeTraits := query.isSelected(formatField(traitField))
	includeTraitRarity := query.isSelected(formatField(traitField, rarityField))
	includeRarity := query.isSelected(formatField(rarityField))

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
func (s *Server) nfts(owner *string, collection *string, rarityMax *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	// FIXME: Change rarity handling here.

	nfts, err := s.storage.NFTs(owner, collection, orderBy)
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

	// Sort NFTs by rarity if needed.
	if orderBy.Field == api.NFTOrderFieldRarity {

		out := make([]*api.NFT, 0, len(nfts))
		for _, nft := range nfts {
			nft := nft

			// FIXME: Retrieve rarity for NFT.
			rarity := 1.0

			// If the NFT is above the rarity threshold, skip it.
			if rarityMax != nil && rarity > *rarityMax {
				continue
			}

			// Include the NFT in the result set.
			out = append(out, nft)
		}

		// FIXME: Better performance can be achieved by inserting to a slice
		// in a way that it remains sorted along the way.
		sort.Slice(out, func(i, j int) bool {
			ri := out[i].Rarity
			rj := out[j].Rarity
			if orderBy.Direction == api.OrderDirectionAsc {
				return ri < rj
			}
			return ri > rj
		})

		nfts = out
	}

	return nfts, nil
}
