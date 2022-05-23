package api

import (
	"context"

	"github.com/NFT-com/analytics/graph/models/api"
	"github.com/NFT-com/analytics/graph/query"
)

// getCollection returns a single collection based on its ID.
func (s *Server) getCollection(ctx context.Context, id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	// Does this query require retrieving the list of NFTs?
	sel := query.GetSelection(ctx)
	includeNFTs := sel.Has(nftField)
	if !includeNFTs {
		return collection, nil
	}

	return s.expandCollectionDetails(ctx, collection)
}

// getCollectionByContract returns a single collection for the specified network, given its contract address.
func (s *Server) getCollectionByContract(ctx context.Context, networkID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByContract(networkID, contract)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Str("contract", contract).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	// Does this query require retrieving the list of NFTs?
	sel := query.GetSelection(ctx)
	includeNFTs := sel.Has(nftField)
	if !includeNFTs {
		return collection, nil
	}

	return s.expandCollectionDetails(ctx, collection)
}

// expandCollectionDetails is the workhorse function that will do all of the heavy lifting for
// the collection queries. It fetches all NFTs from that collection
// (similar to how dataloaders would), but also retrieves traits and deals with rarity calculation.
// NOTE: This function modifies the provided collection in-place.
func (s *Server) expandCollectionDetails(ctx context.Context, collection *api.Collection) (*api.Collection, error) {

	// Retrieve the list of NFTs.
	nfts, err := s.getCollectionNFTs(collection.ID)
	if err != nil {
		return nil, errRetrieveNFTFailed
	}

	s.log.Debug().
		Str("id", collection.ID).
		Int("collection_size", len(nfts)).
		Msg("retrieved list of nfts for collection")

	collection.NFTs = nfts

	// Parse the NFT query.
	cfg := nftQueryConfig{
		traitPath:       query.FieldPath(nftField, traitField),
		traitRarityPath: query.FieldPath(nftField, traitField, rarityField),
		rarityPath:      query.FieldPath(nftField, rarityField),
	}
	req := parseNFTQueryWithConfig(cfg, ctx)

	s.log.Debug().
		Bool("include_nft_rarity", req.nftRarity).
		Bool("include_traits", req.traits).
		Bool("include_trait_rarity", req.traitRarity).
		Msg("NFT information requested")

	// If we do not need traits nor rarity, we're done.
	if !req.traits && !req.needRarity() {
		return collection, nil
	}

	traits, err := s.getTraitsForCollection(collection.ID)
	if err != nil {
		return nil, errRetrieveTraitsFailed
	}

	// Link traits to corresponding NFT.
	for _, nft := range collection.NFTs {
		nft.Traits = traits[nft.ID]
	}

	// If we need traits but not rarity information, just fetch trait information
	// and link them to correct NFTs.
	if !req.needRarity() {
		return collection, nil
	}

	// Crunch the data and determine trait frequency.
	stats := traits.CalculateStats()

	// Total number of NFTs in a collection, in relation to which we're calculating frequency.
	total := len(collection.NFTs)

	// Calculate trait rarity.
	for _, nft := range collection.NFTs {

		rarity, traitRarity := stats.CalculateRarity(uint(total), nft.Traits)

		nft.Rarity = rarity
		// Set this only if individual trait rarity is requested, since it includes
		// traits not necessarily found in this NFT.
		if req.traitRarity {
			nft.Traits = traitRarity
		}
	}

	return collection, nil
}

// getCollectionNFTs returns a list of NFTs in a collection.
func (s *Server) getCollectionNFTs(collectionID string) ([]*api.NFT, error) {

	nfts, err := s.storage.CollectionNFTs(collectionID)
	if err != nil {
		s.logError(err).
			Str("id", collectionID).
			Msg("could not retrieve NFTs for a collection")
		return nil, errRetrieveNFTFailed
	}

	return nfts, nil
}

// collections returns a list of collections according to the specified search criteria and sorting options.
func (s *Server) collections(ctx context.Context, network *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	collections, err := s.storage.Collections(network, orderBy)
	if err != nil {
		log := s.logError(err)
		if network != nil {
			log = log.Str("network", *network)
		}
		log.Msg("could not retrieve collections")
		return nil, errRetrieveCollectionFailed
	}

	for _, collection := range collections {
		collection, err = s.expandCollectionDetails(ctx, collection)
		if err != nil {
			s.logError(err).Str("id", collection.ID).Msg("retrieving collection details failed")
			return nil, errRetrieveCollectionFailed
		}
	}

	return collections, nil
}

// collectionsByNetwork returns a list of collections on a given network.
func (s *Server) collectionsByNetwork(ctx context.Context, networkID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByNetwork(networkID)
	if err != nil {
		s.logError(err).
			Str("network", networkID).
			Msg("could not retrieve collections for a network")
		return nil, errRetrieveCollectionFailed
	}

	for _, collection := range collections {
		collection, err = s.expandCollectionDetails(ctx, collection)
		if err != nil {
			s.logError(err).Str("id", collection.ID).Msg("retrieving collection details failed")
			return nil, errRetrieveCollectionFailed
		}
	}

	return collections, nil
}

// CollectionListings returns a list of marketplaces where the collection is listed on.
func (s *Server) collectionsListings(collectionID string) ([]*api.Marketplace, error) {

	marketplaces, err := s.storage.MarketplacesForCollection(collectionID)
	if err != nil {
		s.logError(err).
			Str("collection", collectionID).
			Msg("could not retrieve marketplaces for a collection")
		return nil, errRetrieveMarketplaceFailed
	}

	return marketplaces, nil
}
