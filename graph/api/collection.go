package api

import (
	"context"

	"github.com/NFT-com/graph-api/graph/models/api"
)

// processCollection is the workhorse function that will do all of the heavy lifting for
// the collection queries. Contrary to the `getCollection` request that only retrieves
// the collection from storage, `processCollection` will (if required) fetch all NFTs
// from that collection (similar to how dataloaders would), but will also retrieve
// required data for rarity calculation.
func (s *Server) processCollection(ctx context.Context, id string) (*api.Collection, error) {

	query := getQuerySelection(ctx)

	// Does this query require retrieving the list of NFTs?
	includeNFTs := query.isSelected(nftField)

	s.log.Debug().
		Str("id", id).
		Bool("include_nfs", includeNFTs).
		Msg("processing collection request")

	collection, err := s.getCollection(id)
	if err != nil {
		return nil, errRetrieveCollectionFailed
	}

	// If we don't need the list of NFTs, we're done.
	if !includeNFTs {
		return collection, nil
	}

	// Retrieve the list of NFTs.
	nfts, err := s.getCollectionNFTs(id)
	if err != nil {
		return nil, errRetrieveNFTFailed
	}

	s.log.Debug().
		Str("id", id).
		Int("collection_size", len(nfts)).
		Msg("retrieved list of collection nfts")

	collection.NFTs = nfts

	includeTraits := query.isSelected(formatField(nftField, traitField))
	includeTraitRarity := query.isSelected(formatField(nftField, traitField, rarityField))
	includeRarity := query.isSelected(formatField(nftField, rarityField))

	s.log.Debug().
		Bool("include_rarity", includeRarity).
		Bool("include_traits", includeTraits).
		Bool("include_trait_rarity", includeTraitRarity).
		Msg("NFT information requested")

	needRarity := includeRarity || includeTraitRarity

	// If we do not need traits nor rarity, we're done.
	if !includeTraits && !needRarity {
		return collection, nil
	}

	traits, err := s.getTraitsForCollection(id)
	if err != nil {
		return nil, errRetrieveTraitsFailed
	}

	// Link traits to corresponding NFT.
	for _, nft := range collection.NFTs {
		nft.Traits = traits[nft.ID]
	}

	// If we need traits but not rarity information, just fetch trait information
	// and link them to correct NFTs.
	if includeTraits && !needRarity {
		return collection, nil
	}

	// Crunch the data and determine trait frequency.
	stats := extractTraitStats(traits)

	stats.Print()

	// Total number of NFTs in a collection, in relation to which we're calculating frequency.
	total := len(collection.NFTs)

	// Calculate trait rarity.
	for _, nft := range collection.NFTs {

		rarity, traitRarity := calcTraitCollectionRarity(uint(total), stats, nft.Traits)

		nft.Rarity = rarity
		// Set this only if individual trait rarity is requested, since it includes
		// traits not necessarily found in this NFT.
		if includeTraitRarity {
			nft.Traits = traitRarity
		}
	}

	return collection, nil
}

// getCollection returns a single collection based on its ID.
func (s *Server) getCollection(id string) (*api.Collection, error) {

	collection, err := s.storage.Collection(id)
	if err != nil {
		s.logError(err).
			Str("id", id).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
	}

	return collection, nil
}

// getCollectionByContract returns a single collection for the specified chain, given its contract address.
func (s *Server) getCollectionByContract(chainID string, contract string) (*api.Collection, error) {

	collection, err := s.storage.CollectionByContract(chainID, contract)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Str("contract", contract).
			Msg("could not retrieve collection")
		return nil, errRetrieveCollectionFailed
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
func (s *Server) collections(chain *string, orderBy api.CollectionOrder) ([]*api.Collection, error) {

	collections, err := s.storage.Collections(chain, orderBy)
	if err != nil {
		log := s.logError(err)
		if chain != nil {
			log = log.Str("chain", *chain)
		}
		log.Msg("could not retrieve collections")
		return nil, errRetrieveCollectionFailed
	}

	return collections, nil
}

// collectionsByChain returns a list of collections on a given chain.
func (s *Server) collectionsByChain(chainID string) ([]*api.Collection, error) {

	collections, err := s.storage.CollectionsByChain(chainID)
	if err != nil {
		s.logError(err).
			Str("chain", chainID).
			Msg("could not retrieve collections for a chain")
		return nil, errRetrieveCollectionFailed
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
