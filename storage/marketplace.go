package storage

import (
	"fmt"

	"github.com/NFT-com/indexer-api/models/api"
)

// MarketplacesForCollection returns all Marketplaces that the specified Collection is associated with.
func (s *Storage) MarketplacesForCollection(collectionID string) ([]*api.Marketplace, error) {

	var marketplaces []*api.Marketplace

	err := s.db.Debug().
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
