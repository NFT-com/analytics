package storage

import (
	"github.com/NFT-com/indexer-api/models/api"
)

var firstNft = api.Nft{
	ID:         "id-1",
	TokenID:    14,
	Owner:      "0xAC0",
	URI:        "dummy-uri",
	Rarity:     14.0,
	Collection: &firstCollection,
}

var secondNft = api.Nft{
	ID:         "id-2",
	TokenID:    15,
	Owner:      "0xAC0",
	URI:        "dummy-uri-2",
	Rarity:     140.0,
	Collection: &secondCollection,
}

var sampleNfts = []*api.Nft{
	&firstNft,
	&secondNft,
}

var chain = api.Chain{
	ID:          "chain-1",
	Name:        "Ethereum",
	Description: "Description of an Ethereum.",
}

var firstCollection = api.Collection{
	ID:          "col-id-1",
	Name:        "Distracted Ape Sailing Group",
	Description: "I still think this is a fad",
	Address:     "0xAC0ADD1",
	Chain:       &chain,
}

var secondCollection = api.Collection{
	ID:          "col-id-2",
	Name:        "CryptoMisfits",
	Description: "I still think this is a fad",
	Address:     "0xAC0ADD2",
	Chain:       &chain,
}

var sampleCollections = []*api.Collection{
	&firstCollection,
	&secondCollection,
}
