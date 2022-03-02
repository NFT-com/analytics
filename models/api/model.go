// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"strconv"
)

// Chain represents the chain and its networks.
type Chain struct {
	// Chain ID.
	ID string `json:"id"`
	// Name of the chain, e.g. `Ethereum`.
	Name string `json:"name"`
	// Description of the chain.
	Description string `json:"description"`
	// Marketplaces on this chain.
	Marketplaces []*Marketplace `json:"marketplaces"`
	// Collections found on this chain.
	Collections []*Collection `json:"collections"`
}

// Collection represents a group of NFTs that share the same smart contract.
type Collection struct {
	// Collection ID.
	ID string `json:"id"`
	// Collection name, e.g. `CryptoKitties`.
	Name string `json:"name"`
	// Description of the collection.
	Description string `json:"description"`
	// Address of the smart-contract.
	Address string `json:"address"`
	// Chain on which collection resides on.
	Chain *Chain `json:"chain"`
	// Marketplaces this collection is on.
	Marketplaces []*Marketplace `json:"marketplaces"`
	// List of NFTs that are part of this collection.
	Nfts []*Nft `json:"nfts"`
}

// Ordering options for collections.
type CollectionOrder struct {
	// Field by which collections should be sorted.
	Field CollectionOrderField `json:"field"`
	// Direction in which collections should be sorted.
	Direction OrderDirection `json:"direction"`
}

// Marketplace represents a single NFT marketplace (e.g. Opensea, DefiKingdoms).
type Marketplace struct {
	// Marketplace ID.
	ID string `json:"id"`
	// Marketplace name, e.g. `Opensea`.
	Name string `json:"name"`
	// Description of the marketplace.
	Description string `json:"description"`
	// Chains the marketplace operates on.
	Chains []*Chain `json:"chains"`
	// Collections on this marketplace.
	Collections []*Collection `json:"collections"`
}

type Nft struct {
	// NFT ID.
	ID string `json:"id"`
	// Token ID, as found on the blockchain.
	TokenID int `json:"token_id"`
	// Address of the account that owns the NFT.
	Owner string `json:"owner"`
	// URI of the NFT, directing to e.g. a JSON file with asset metadata.
	URI string `json:"uri"`
	// Rarity score for the NFT.
	Rarity float64 `json:"rarity"`
	// Collection this NFT is part of.
	Collection *Collection `json:"collection"`
}

// Ordering options for NFTs.
type NFTOrder struct {
	// Field by which NFTs should be sorted by.
	Field NFTOrderField `json:"field"`
	// Direction in which NFTs should be sorted.
	Direction OrderDirection `json:"direction"`
}

// Properties by which collections can be ordered.
type CollectionOrderField string

const (
	// Order by creation time.
	CollectionOrderFieldCreationTime CollectionOrderField = "CREATION_TIME"
	// Order by market cap.
	CollectionOrderFieldMarketCap CollectionOrderField = "MARKET_CAP"
	// Order by total volume.
	CollectionOrderFieldTotalVolume CollectionOrderField = "TOTAL_VOLUME"
	// Order by biggest gains.
	CollectionOrderFieldBiggestGains CollectionOrderField = "BIGGEST_GAINS"
	// Order by biggest losses.
	CollectionOrderFieldBiggestLosses CollectionOrderField = "BIGGEST_LOSSES"
	// Order by daily volume.
	CollectionOrderFieldDailyVolume CollectionOrderField = "DAILY_VOLUME"
)

var AllCollectionOrderField = []CollectionOrderField{
	CollectionOrderFieldCreationTime,
	CollectionOrderFieldMarketCap,
	CollectionOrderFieldTotalVolume,
	CollectionOrderFieldBiggestGains,
	CollectionOrderFieldBiggestLosses,
	CollectionOrderFieldDailyVolume,
}

func (e CollectionOrderField) IsValid() bool {
	switch e {
	case CollectionOrderFieldCreationTime, CollectionOrderFieldMarketCap, CollectionOrderFieldTotalVolume, CollectionOrderFieldBiggestGains, CollectionOrderFieldBiggestLosses, CollectionOrderFieldDailyVolume:
		return true
	}
	return false
}

func (e CollectionOrderField) String() string {
	return string(e)
}

func (e *CollectionOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CollectionOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CollectionOrderField", str)
	}
	return nil
}

func (e CollectionOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which NFTs could be ordered by.
type NFTOrderField string

const (
	// Order by creation time.
	NFTOrderFieldCreationTime NFTOrderField = "CREATION_TIME"
	// Order by rarity.
	NFTOrderFieldRarity NFTOrderField = "RARITY"
	// Order by value.
	NFTOrderFieldValue NFTOrderField = "VALUE"
)

var AllNFTOrderField = []NFTOrderField{
	NFTOrderFieldCreationTime,
	NFTOrderFieldRarity,
	NFTOrderFieldValue,
}

func (e NFTOrderField) IsValid() bool {
	switch e {
	case NFTOrderFieldCreationTime, NFTOrderFieldRarity, NFTOrderFieldValue:
		return true
	}
	return false
}

func (e NFTOrderField) String() string {
	return string(e)
}

func (e *NFTOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NFTOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NFTOrderField", str)
	}
	return nil
}

func (e NFTOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Available options for the `orderBy` direction argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a decending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
