package storage

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// MarketplacesForCollection returns all Marketplaces that the specified Collection is associated with.
func (s *Storage) MarketplacesForCollection(collectionID string) ([]*api.Marketplace, error) {

	var marketplaces []*api.Marketplace

	err := s.db.
		Table("collection c").
		Select("m.*").
		Joins("INNER JOIN marketplace_collections mc ON c.id = mc.collection_id").
		Joins("INNER JOIN marketplace m ON mc.marketplace_id = m.id").
		Where("c.id = ?", collectionID).
		Find(&marketplaces).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceCollections returns all Collections associated with the specified Marketplace.
func (s *Storage) MarketplaceCollections(marketplaceID string) ([]*api.Collection, error) {

	var collections []*api.Collection

	err := s.db.
		Table("marketplace m").
		Select("c.*").
		Joins("INNER JOIN marketplace_collections mc ON m.id = mc.marketplace_id").
		Joins("INNER JOIN collection c ON mc.collection_id = c.id").
		Where("m.id = ?", marketplaceID).
		Find(&collections).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collections: %w", err)
	}

	return collections, nil
}

// Retrieve a list of Marketplaces on a specified Chain.
func (s *Storage) MarketplacesByChain(chainID string) ([]*api.Marketplace, error) {

	// FIXME: Does it make a difference to use `Table("marketplace m")` vs `Table("collection c")` ?
	// If not, make it uniform.

	var marketplaces []*api.Marketplace

	// FIXME: Suspect using IN would be faster here.

	err := s.db.
		Table("marketplace m").
		Select("DISTINCT m.*").
		Joins("INNER JOIN marketplace_collections mc ON m.id = mc.marketplace_id").
		Joins("INNER JOIN collection c ON mc.collection_id = c.id").
		Where("c.chain_id = ?", chainID).
		Find(&marketplaces).
		Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve marketplaces: %w", err)
	}

	return marketplaces, nil
}

// MarketplaceChains retrieves all chains that the Marketplace supports.
func (s *Storage) MarketplaceChains(marketplaceID string) ([]*api.Chain, error) {

	var chains []*api.Chain

	selection := s.db.
		Select("DISTINCT c.chain_id").
		Table("marketplace_collections mc").
		Joins("INNER JOIN collection c ON mc.collection_id = c.id").
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
