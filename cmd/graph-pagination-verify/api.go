package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/NFT-com/analytics/graph/models/api"
)

const (
	fmtCollectionQuery = `
query {
	collection(id: "%s") {
		id
		nfts(first: %d, after: "%s") {
		edges {
			node {
			id
			}
			cursor
		}
		pageInfo {
			hasNextPage
		}
		}
	}
}`
)

// getPageFromAPI retrieves the NFTs in a collection, according the the specified paramters.
func getPageFromAPI(graphAPI string, collectionID string, dump io.Writer, first uint, after string) (api.NFTConnection, error) {

	query := fmt.Sprintf(fmtCollectionQuery, collectionID, first, after)

	// Prepase the request.
	request := graphQLQuery{
		Query: query,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return api.NFTConnection{}, fmt.Errorf("could not encode API request: %w", err)
	}

	// Execute the request.
	res, err := http.Post(graphAPI, "application/json", bytes.NewReader(data))
	if err != nil {
		return api.NFTConnection{}, fmt.Errorf("could not execute POST request: %w", err)
	}
	defer res.Body.Close()

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return api.NFTConnection{}, fmt.Errorf("could not read response data: %w", err)
	}

	key := fmt.Sprintf("%v-%v-%v", collectionID, first, after)
	err = dumpPayload(dump, key, payload)
	if err != nil {
		return api.NFTConnection{}, fmt.Errorf("could not log response data: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return api.NFTConnection{}, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	var response apiResponse
	err = json.Unmarshal(payload, &response)
	if err != nil {
		return api.NFTConnection{}, fmt.Errorf("could not unmarshal response data: %w", err)
	}

	return response.Data.Collection.NFTs, nil
}
