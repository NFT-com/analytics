package api

import (
	"fmt"
	"sort"

	"github.com/NFT-com/graph-api/graph/models/api"
)

// getNFT returns a single NFT based on its ID.
func (s *Server) getNFT(id string) (*api.NFT, error) {

	nft, err := s.storage.NFT(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return nft, nil
}

// getNFTByTokenID returns a single NFT based on the combination of chainID, contract address and token ID.
func (s *Server) getNFTByTokenID(chainID string, contract string, tokenID string) (*api.NFT, error) {

	nft, err := s.storage.NFTByTokenID(chainID, contract, tokenID)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Str("contract", contract).
			Str("token_id", tokenID).
			Msg("could not retrieve nft")
		return nil, errRetrieveNFTFailed
	}

	return nft, nil
}

// nfts returns a list of NFTs fitting the search criteria.
func (s *Server) nfts(owner *string, collection *string, rarityMin *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	nfts, err := s.storage.NFTs(owner, collection, orderBy)
	if err != nil {
		log := s.logError(err)
		if owner != nil {
			log = log.Str("owner", *owner)
		}
		if collection != nil {
			log = log.Str("collection", *collection)
		}
		if rarityMin != nil {
			log = log.Float64("min_rarity", *rarityMin)
		}
		log.Msg("could not retrieve nfts")
		return nil, errRetrieveNFTFailed
	}

	// Sort NFTs by rarity if needed.
	if orderBy.Field == api.NFTOrderFieldRarity {

		out := make([]*api.NFT, 0, len(nfts))
		rarities := make([]float64, 0, len(nfts))

		for _, nft := range nfts {
			nft := nft

			// Get rarity information for individual traits.
			traits, err := s.nftTraits(nft.ID)
			if err != nil {
				return nil, fmt.Errorf("could not retrieve trait for an NFT: %w", err)
			}

			// Get rarity for the NFT as a whole.
			rarity := calcRarity(traits)

			// If the NFT is below the rarity threshold, skip it.
			if rarityMin != nil && rarity < *rarityMin {
				continue
			}

			// Include the NFT in the result set.
			out = append(out, nft)
			rarities = append(rarities, rarity)
		}

		// FIXME: Better performance can be achieved by inserting to a slice
		// in a way that it remains sorted along the way.
		sort.Slice(out, func(i, j int) bool {
			return rarities[i] < rarities[j]
		})

		nfts = out
	}

	return nfts, nil
}

// calcRarity calculates the rarity of an NFT by multiplying the
// ratios of individual traits.
// For example, if NFT has two traits that are present in 50% of
// NFTs in a collection, the rarity will be calculated as 0.5*0.5 = 0.25.
func calcRarity(traits []*api.TraitRatio) float64 {

	rarity := 1.0
	for _, trait := range traits {
		rarity = rarity * trait.Ratio
	}

	return rarity
}
