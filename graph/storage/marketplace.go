package storage

import (
	"fmt"

	"github.com/NFT-com/analytics/graph/models/api"
)

// TODO: Refactor the queries below to not use explicit joins but leave it to the query compiler
// to determine the best course of action.
// See https://github.com/NFT-com/analytics/issues/6

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

// MarketplacesByNetwork retrieves a list of marketplaces on a specified network.
func (s *Storage) MarketplacesByNetwork(networkID string) ([]*api.Marketplace, error) {

	var marketplaces []*api.Marketplace
	err := s.db.
		Table("marketplaces m, networks_marketplaces nm").
		Select("DISTINCT m.*").
		Where("m.id = nm.marketplace_id").
		Where("nm.network_id = ?", networkID).
		Find(&marketplaces).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceNetworks retrieves all networks that the marketplace supports.
func (s *Storage) MarketplaceNetworks(marketplaceID string) ([]*api.Network, error) {

	var networks []*api.Network
	selection := s.db.
		Table("marketplaces_collections mc").
		Select("DISTINCT c.network_id").
		Joins("INNER JOIN collections c ON mc.collection_id = c.id").
		Where("mc.marketplace_id = ?", marketplaceID)

	err := s.db.
		Where("id IN (?)", selection).
		Find(&networks).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve networks: %w", err)
	}

	return networks, nil
}
