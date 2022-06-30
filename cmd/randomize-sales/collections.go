package main

import (
	"fmt"

	"gorm.io/gorm"
)

type collectionRecord struct {
	ID      string `gorm:"column:id"`
	Address string `gorm:"column:contract_address"`
}

func getCollections(db *gorm.DB) ([]collectionRecord, error) {

	var collections []collectionRecord
	err := db.Raw("SELECT id, contract_address FROM collections").Scan(&collections).Error
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection IDs: %s", err)
	}

	return collections, nil
}
