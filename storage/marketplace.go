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

// MarketplaceCollectionsList returns all Collections associated with the specified Marketplace.
func (s *Storage) MarketplaceCollectionsList(marketplaceID string) ([]*api.Collection, error) {

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
