package storage

import (
	"fmt"

	"github.com/NFT-com/graph-api/models/api"
)

// TODO: Refactor the queries below to not use explicit joins but leave it to the query compiler
// to determine the best course of action.
// See https://github.com/NFT-com/graph-api/issues/6

// MarketplacesForCollection retrieves all marketplaces that the specified collection is associated with.
func (s *Storage) MarketplacesForCollection(collectionID string) ([]*api.Marketplace, error) {

	var marketplaces []*api.Marketplace

	err := s.db.
		Table("marketplaces m").
		Select("m.*").
		Joins("INNER JOIN marketplaces_collections mc ON m.id = mc.marketplace_id").
		Joins("INNER JOIN collections c ON mc.collection_id = c.id").
		Where("c.id = ?", collectionID).
		Find(&marketplaces).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceCollections retrieves all collections associated with the specified marketplace.
func (s *Storage) MarketplaceCollections(marketplaceID string) ([]*api.Collection, error) {

	var collections []*api.Collection

	err := s.db.
		Table("marketplaces m").
		Select("c.*").
		Joins("INNER JOIN marketplaces_collections mc ON m.id = mc.marketplace_id").
		Joins("INNER JOIN collections c ON mc.collection_id = c.id").
		Where("m.id = ?", marketplaceID).
		Find(&collections).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// MarketplacesByChain retrieves a list of marketplaces on a specified Chain.
func (s *Storage) MarketplacesByChain(chainID string) ([]*api.Marketplace, error) {

	var marketplaces []*api.Marketplace

	err := s.db.
		Table("marketplaces m").
		Select("DISTINCT m.*").
		Joins("INNER JOIN marketplaces_collections mc ON m.id = mc.marketplace_id").
		Joins("INNER JOIN collections c ON mc.collection_id = c.id").
		Where("c.chain_id = ?", chainID).
		Find(&marketplaces).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceChains retrieves all chains that the marketplace supports.
func (s *Storage) MarketplaceChains(marketplaceID string) ([]*api.Chain, error) {

	var chains []*api.Chain

	selection := s.db.
		Table("marketplaces_collections mc").
		Select("DISTINCT c.chain_id").
		Joins("INNER JOIN collections c ON mc.collection_id = c.id").
		Where("mc.marketplace_id = ?", marketplaceID)

	err := s.db.
		Where("id IN (?)", selection).
		Find(&chains).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve chains: %w", err)
	}

	return chains, nil
}
