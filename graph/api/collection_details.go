package api

import (
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/NFT-com/analytics/graph/api/internal/query"
	"github.com/NFT-com/analytics/graph/models/api"
)

// expandCollectionDetails adds the NFT information to the collection object,
// as well as the needed collection stats.
func (s *Server) expandCollectionDetails(query *query.Collection, collection *api.Collection) error {

	// Retrieve any statistics from the aggregation API.
	err := s.expandCollectionStats(query, collection)
	if err != nil {
		// Continue even if stats could not be retrieved (e.g. API was unavailable).
		s.log.Error().Err(err).Str("id", collection.ID).Msg("could not retrieve collection stats")
	}

	// If we don't need any NFT data, we're done.
	if !query.NFTs {
		return nil
	}

	err = s.expandCollectionNFTData(query, collection)
	if err != nil {
		return fmt.Errorf("could not get nft data: %w", err)
	}

	return nil
}

// expandCollectionStats retrieves the collection stats from the aggregation API.
func (s *Server) expandCollectionStats(query *query.Collection, collection *api.Collection) error {

	// Execute as much as possible, return the composite error in the end.
	var multiErr error

	// Get volume from the aggregation API.
	if query.Volume {
		volumes, err := s.aggregationAPI.CollectionVolumes([]string{collection.ID})
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get collection volume: %w", err))
		}

		collection.Volume = volumes[collection.ID]
	}

	// Get market cap from the aggregation API.
	if query.MarketCap {
		caps, err := s.aggregationAPI.CollectionMarketCaps([]string{collection.ID})
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get collection market cap: %w", err))
		}

		collection.MarketCap = caps[collection.ID]
	}

	// Get sale count from the aggregation API.
	if query.Sales {
		sales, err := s.aggregationAPI.CollectionSales(collection.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get collection sales: %w", err))
		}

		collection.Sales = uint64(sales)
	}

	return multiErr
}

// expandCollectionNFTData is the workhorse function that will do all of the heavy lifting for
// the collection queries. It fetches all NFTs from that collection
// (similar to how dataloaders would), but also retrieves traits and deals with rarity calculation.
// NOTE: This function modifies the provided collection in-place.
func (s *Server) expandCollectionNFTData(query *query.Collection, collection *api.Collection) error {

	// Retrieve the list of NFTs.
	nfts, err := s.getCollectionNFTs(collection.ID)
	if err != nil {
		return fmt.Errorf("could not get NFTs in the collection: %w", err)
	}

	s.log.Debug().
		Str("id", collection.ID).
		Int("collection_size", len(nfts)).
		Msg("retrieved list of nfts for collection")

	collection.NFTs = nfts

	s.log.Debug().
		Bool("rarity", query.NFT.Fields.Rarity).
		Bool("traits", query.NFT.Fields.Traits).
		Bool("trait_rarity", query.NFT.Fields.TraitRarity).
		Bool("owners", query.NFT.Fields.Owners).
		Bool("price", query.NFT.Fields.Price).
		Bool("average_price", query.NFT.Fields.AveragePrice).
		Msg("NFT information requested")

	// Retrieve owners if needed.
	if query.NFT.Fields.Owners {
		owners, err := s.storage.CollectionOwners(collection.ID)
		if err != nil {
			return fmt.Errorf("could not retrieve owners for the collection: %w", err)
		}

		// Set the owners for each of the NFTs.
		for _, nft := range collection.NFTs {
			nft.Owners = owners[nft.ID]
		}
	}

	// If NFT prices are required, retrieve them now.
	if query.NFT.Fields.Price || query.NFT.Fields.AveragePrice {

		// Retrieve stats, but continue even if some could not be retrieved (e.g. API was unavailable).
		var prices map[string]float64
		if query.NFT.Fields.Price {
			prices, err = s.aggregationAPI.CollectionPrices(collection.ID)
			if err != nil {
				s.log.Error().Err(err).Msg("could not retrieve NFT prices")
			}
		}
		var averages map[string]float64
		if query.NFT.Fields.AveragePrice {
			averages, err = s.aggregationAPI.CollectionAveragePrices(collection.ID)
			if err != nil {
				s.log.Error().Err(err).Msg("could not retrieve NFT average prices")
			}
		}

		// Set the appropriate price fields.
		for _, nft := range collection.NFTs {
			nft.TradingPrice = prices[nft.ID]
			nft.AveragePrice = averages[nft.ID]
		}
	}

	// If we do not need traits nor rarity, we're done.
	if !query.NFT.Fields.Traits && !query.NFT.Fields.NeedRarity() {
		return nil
	}

	// Fetch traits for this collection.
	traits, err := s.getTraitsForCollection(collection.ID)
	if err != nil {
		return fmt.Errorf("could not get traits for the collection: %w", err)
	}

	// Link traits to corresponding NFT.
	for _, nft := range collection.NFTs {
		nft.Traits = traits[nft.ID]
	}

	// If don't need rarity information, we're done.
	if !query.NFT.Fields.NeedRarity() {
		return nil
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
		if query.NFT.Fields.TraitRarity {
			nft.Traits = traitRarity
		}
	}

	return nil
}
