package main

import (
	"encoding/json"

	"github.com/NFT-com/analytics/graph/models/api"
)

type apiResponse struct {
	Data responseData `json:"data"`
}

type responseData struct {
	Collection collectionData `json:"collection"`
}

type collectionData struct {
	ID   string            `json:"id"`
	NFTs api.NFTConnection `json:"nfts"`
}

type graphQLQuery struct {
	Query string `json:"query"`
}

type responseDump struct {
	Key  string          `json:"key"`
	Data json.RawMessage `json:"data"`
}
