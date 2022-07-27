package datapoint

import (
	"time"
)

// Value represents the generic datatype for some stat.
type Value struct {
	ID    string  `json:"id,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// Volume represents the trading volume for a given category (e.g. collection or a marketplace).
type Volume struct {
	Total float64    `json:"total" gorm:"column:total"`
	Date  *time.Time `json:"date,omitempty" gorm:"column:date"`
}

// MarketCap represents the total market cap of a given entity (collection or a marketplace).
type MarketCap struct {
	Total float64    `json:"total" gorm:"column:total"`
	Date  *time.Time `json:"date,omitempty" gorm:"column:date"`
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
	LowestRecentPrice float64 `json:"lowest_price" gorm:"column:lowest_price"`
	Start             string  `json:"start" gorm:"column:start_date"`
	End               string  `json:"end" gorm:"column:end_date"`
}

// Average represents the average price of an NFT in a collection.
type Average struct {
	Average float64    `json:"average" gorm:"column:average"`
	Date    *time.Time `json:"date,omitempty" gorm:"column:date"`
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
	Price float64    `json:"price" gorm:"column:trade_price"`
	Time  *time.Time `json:"emitted_at,omitempty" gorm:"column:emitted_at"`
}

// Count represents the total number of NFTs in a collection.
type Count struct {
	Mints uint64     `json:"mints" gorm:"column:mints"`
	Burns uint64     `json:"burns" gorm:"column:burns"`
	Date  *time.Time `json:"date,omitempty" gorm:"column:date"`
}
