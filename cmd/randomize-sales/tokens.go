package main

import (
	"fmt"

	"gorm.io/gorm"
)

// getTokenMap returns tokenIDs, mapped to the collection contract addresses.
func getTokenMap(db *gorm.DB, collections []collectionRecord) (map[string][]string, error) {

	tokens := make(map[string][]string)

	for _, collection := range collections {

		var ids []string
		err := db.Raw("SELECT token_id FROM nfts WHERE collection_id = ? ORDER BY token_id ASC", collection.ID).Scan(&ids).Error
		if err != nil {
			return nil, fmt.Errorf("could not retrieve token IDs: %w", err)
		}

		tokens[collection.Address] = ids
	}

	return tokens, nil
}
