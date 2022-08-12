package datapoint

import (
	"time"
)

// Count represents a generic datatype for some count-related stat, e.g. a user or sale count.
type Count struct {
	ID    string `json:"id,omitempty"`
	Value uint64 `json:"value,omitempty"`
}

// Sale represents the number of sales for a given category (e.g. collection or a marketplace).
type Sale struct {
	Count uint64     `json:"count" gorm:"column:count"`
	Date  *time.Time `json:"date,omitempty" gorm:"column:date"`
}

// LowestPrice represents the minimum price for an NFT in a collection.
// `Start` and `End` values denote the time span over which the
// minimum trading price is calculated.
type LowestPrice struct {
	Currency Currency `json:"lowest_price"`
	Start    string   `json:"start"`
	End      string   `json:"end"`
}

// Users represents the number of unique users on a marketplace.
type Users struct {
	Count uint64     `json:"count" gorm:"column:count"`
	Date  *time.Time `json:"date,omitempty" gorm:"column:date"`
}

// Price represents the price of an NFT.
// NOTE: This is the only data type that uses the actual time instead
// of the date.
type Price struct {
	Currency Currency   `json:"price"`
	Time     *time.Time `json:"emitted_at,omitempty"`
}

// CollectionSize represents the total number of NFTs in a collection.
type CollectionSize struct {
	Mints uint64     `json:"mints" gorm:"column:mints"`
	Burns uint64     `json:"burns" gorm:"column:burns"`
	Date  *time.Time `json:"date,omitempty" gorm:"column:date"`
}
