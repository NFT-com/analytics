package api

import (
	"context"
	"fmt"

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

	// Parse the query to know how much information to return/calculate.
	query := parseNFTQuery(ctx)

	// Get NFT details.
	err = s.expandNFTDetails(query, nft)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT details: %w", err)
	}

	return nft, nil
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

	// Parse the query to know how much information to return/calculate.
	query := parseNFTQuery(ctx)

	// Get NFT details.
	err = s.expandNFTDetails(query, nft)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT details: %w", err)
	}

	return nft, nil
}

// nfts returns a list of NFTs fitting the search criteria.
func (s *Server) nfts(ctx context.Context, owner *string, collectionID *string, rarityMax *float64, orderBy api.NFTOrder) ([]*api.NFT, error) {

	// Parse the query to know how much information to return/calculate.
	query := parseNFTQuery(ctx)

	nfts, err := s.storage.NFTs(owner, collectionID, orderBy, s.searchLimit, query.Owners)
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

	filterByRarity := rarityMax != nil

	// If we do not need traits nor rarity, we're done.
	if !query.Traits && !query.NeedRarity() && !filterByRarity {
		return nfts, nil
	}

	// Retrieve traits for relevant collections.
	// NOTE: We potentially iterate through the result set twice, first to
	// create the traits map for collections, and then a second time to crunch the numbers and
	// calculate rarity

	// Map collection ID to collection traits.
	traits := make(map[string]collection.TraitMap)

	for _, nft := range nfts {
		// Lookup traits for this collection.
		// If we don't have them already cached, fetch them now.
		ctraits, ok := traits[nft.Collection]
		if !ok {
			tc, err := s.getTraitsForCollection(nft.Collection)
			if err != nil {
				return nil, fmt.Errorf("could not retrieve traits for collection (id: %s): %w", nft.Collection, err)
			}
			traits[nft.Collection] = tc
			ctraits = tc
		}

		// Set the traits for this NFT.
		nft.Traits = ctraits[nft.ID]
	}

	// If we don't need rarity information - we have everything we need and we're done.
	if !query.NeedRarity() && !filterByRarity {
		return nfts, nil
	}

	// We also need to calculate rarity.

	// Cache traits and stats for each of the collections.
	stats := make(map[string]collection.Stats)
	sizes := make(map[string]uint)

	// List of NFTs that fit into the rarity filter specified.
	var filteredNFTs []*api.NFT

	for _, nft := range nfts {
		nft := nft
		// Lookup stats for this collection.
		// If we don't have them already cached, fetch them now, calculate stats and cache them.
		cstats, ok := stats[nft.Collection]
		if !ok {
			tc := traits[nft.Collection]

			size, err := s.storage.CollectionSize(nft.Collection)
			if err != nil {
				s.logError(err).Str("collection", nft.Collection).Msg("could not retrieve collection size")
				return nil, errRetrieveTraitsFailed
			}

			st := tc.CalculateStats()
			stats[nft.Collection] = st
			sizes[nft.Collection] = size

			cstats = st
		}

		// Calculate rarity.
		rarity, traitRarity := cstats.CalculateRarity(sizes[nft.Collection], nft.Traits)
		nft.Rarity = rarity
		// Only set traits rarity if explicitely requested.
		if query.TraitRarity {
			nft.Traits = traitRarity
		}

		// If we're not filtering by rarity just include this NFT.
		if !filterByRarity {
			filteredNFTs = append(filteredNFTs, nft)
			continue
		}

		// If we are filtering NFTs by rarity, check whether this NFT is rare enough.
		if nft.Rarity < *rarityMax {
			filteredNFTs = append(filteredNFTs, nft)
		}
	}

	return filteredNFTs, nil
}
