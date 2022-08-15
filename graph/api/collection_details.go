package api

import (
	"fmt"

	"github.com/hashicorp/go-multierror"

	aggregate "github.com/NFT-com/analytics/aggregate/models/api"
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
		} else {

			// Translate the Aggregation API format to the expected Graph format.
			formatted, err := s.createCurrencyList(volumes[collection.ID])
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf("could not transform currency list for volume: %w", err))
			}

			collection.Volume = formatted
		}
	}

	// Get market cap from the aggregation API.
	if query.MarketCap {
		caps, err := s.aggregationAPI.CollectionMarketCaps([]string{collection.ID})
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get collection market cap: %w", err))
		} else {

			// Translate the Aggregation API format to the expected Graph format.
			formatted, err := s.createCurrencyList(caps[collection.ID])
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf("could not transform currency list for market cap: %w", err))
			}

			collection.Volume = formatted
		}
	}

	// Get sale count from the aggregation API.
	if query.Sales {
		sales, err := s.aggregationAPI.CollectionSales(collection.ID)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("could not get collection sales: %w", err))
		}

		collection.Sales = sales
	}

	return multiErr
}

// expandCollectionNFTData is the workhorse function that will do all of the heavy lifting for
// the collection queries. It fetches all NFTs from that collection
// (similar to how dataloaders would), but also retrieves traits and deals with rarity calculation.
// NOTE: This function modifies the provided collection in-place.
func (s *Server) expandCollectionNFTData(query *query.Collection, collection *api.Collection) error {

	// Determine from which NFT ID we should continue pagination.
	afterID, err := decodeCursor(query.NFT.Arguments.After)
	if err != nil {
		return fmt.Errorf("could not decode pagination cursor: %w", err)
	}

	// Retrieve the list of NFTs.
	nfts, haveMore, err := s.getCollectionNFTs(collection.ID, query.NFT.Arguments.First, afterID)
	if err != nil {
		return fmt.Errorf("could not get NFTs in the collection: %w", err)
	}

	s.log.Debug().
		Str("id", collection.ID).
		Int("count", len(nfts)).
		Str("after", afterID).
		Uint("first", query.NFT.Arguments.First).
		Msg("retrieved list of nfts for collection")

	// Create edge types.
	edges := make([]api.NFTEdge, 0, len(nfts))
	for _, nft := range nfts {
		nft := nft

		edge := api.NFTEdge{
			Node:   nft,
			Cursor: createCursor(nft.ID),
		}

		edges = append(edges, edge)
	}

	pageInfo := api.PageInfo{
		HasNextPage: haveMore,
	}

	// Set start cursor if needed.
	if query.StartCursor {
		firstID, err := s.getFirstID(collection.ID)
		if err != nil {
			return fmt.Errorf("could not retrieve NFT ID for start cursor: %w", err)
		}
		pageInfo.StartCursor = createCursor(firstID)
	}

	nftConn := api.NFTConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}

	collection.NFTs = nftConn

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
		for _, edge := range collection.NFTs.Edges {
			edge.Node.Owners = owners[edge.Node.ID]
		}
	}

	// If NFT prices are required, retrieve them now.
	if query.NFT.Fields.Price || query.NFT.Fields.AveragePrice {

		// Retrieve stats, but continue even if some could not be retrieved (e.g. API was unavailable).
		var prices map[string][]aggregate.Coin
		if query.NFT.Fields.Price {
			prices, err = s.aggregationAPI.CollectionPrices(collection.ID)
			if err != nil {
				s.log.Error().Err(err).Msg("could not retrieve NFT prices")
			}
		}
		var averages map[string][]aggregate.Coin
		if query.NFT.Fields.AveragePrice {
			averages, err = s.aggregationAPI.CollectionAveragePrices(collection.ID)
			if err != nil {
				s.log.Error().Err(err).Msg("could not retrieve NFT average prices")
			}
		}

		for _, edge := range collection.NFTs.Edges {

			price, err := s.createCurrencyList(prices[edge.Node.ID])
			if err != nil {
				// FIXME: See how you logged this elsewhere.
				s.log.Error().Err(err).Msg("could not create currency list for NFT price")
			}

			avg, err := s.createCurrencyList(averages[edge.Node.ID])
			if err != nil {
				// FIXME: See how you logged this elsewhere.
				s.log.Error().Err(err).Msg("could not create currency list for NFT average price")
			}

			edge.Node.TradingPrice = price
			edge.Node.AveragePrice = avg
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
	for _, edge := range collection.NFTs.Edges {
		edge.Node.Traits = traits[edge.Node.ID]
	}

	// If don't need rarity information, we're done.
	if !query.NFT.Fields.NeedRarity() {
		return nil
	}

	// Total number of NFTs in a collection, in relation to which we're calculating frequency.
	total, err := s.storage.CollectionSize(collection.ID)
	if err != nil {
		return fmt.Errorf("could not get collection size: %w", err)
	}

	// Crunch the data and determine trait frequency.
	stats := traits.CalculateStats()

	// Calculate trait rarity.
	for _, edge := range collection.NFTs.Edges {

		rarity, traitRarity := stats.CalculateRarity(total, edge.Node.Traits)

		edge.Node.Rarity = rarity
		// Set this only if individual trait rarity is requested, since it includes
		// traits not necessarily found in this NFT.
		if query.NFT.Fields.TraitRarity {
			edge.Node.Traits = traitRarity
		}
	}

	return nil
}

// getFirstID returns the ID of the first NFT in the collection, when sorted by ID ascending.
func (s *Server) getFirstID(collectionID string) (string, error) {

	nfts, _, err := s.getCollectionNFTs(collectionID, 1, "")
	if err != nil {
		return "", fmt.Errorf("could not get first NFT in the collection: %w", err)
	}

	return nfts[0].ID, nil
}
