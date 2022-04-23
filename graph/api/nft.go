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
func (s *Server) nfts(owner *string, collection *string, rarityMax *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

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

			rarity, cached := nft.GetCachedRarity()
			if !cached {
				// Get trait information.
				traits, err := s.nftTraits(nft)
				if err != nil {
					return nil, fmt.Errorf("could not retrieve traits for NFT: %w", err)
				}

				// Cache the trait/rarity information for potential later use.
				nft.CacheTraits(traits)

				rarity, _ = nft.GetCachedRarity()
			}

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
			ri, _ := out[i].GetCachedRarity()
			rj, _ := out[j].GetCachedRarity()
			if orderBy.Direction == api.OrderDirectionAsc {
				return ri < rj
			}
			return ri > rj
		})

		nfts = out
	}

	return nfts, nil
}
