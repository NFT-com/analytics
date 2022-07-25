package main

import (
	"fmt"

	"gorm.io/gorm"
)

// getCollectionNFTs returns the sorted list IDs of NFTs in the collection.
func getCollectionNFTs(db *gorm.DB, collectionID string) ([]string, error) {

	query := db.
		Table("nfts").
		Select("id").
		Where("collection_id = ?", collectionID).
		Order("id ASC")

	var ids []string
	err := query.Find(&ids).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve NFT ids: %w", err)
	}

	return ids, nil
}
