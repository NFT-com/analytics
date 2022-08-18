package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	contentTypeJSON = "application/json"
)

// POST executes the HTTP request and unpacks the response. Request and responses are interpreted as JSON.
func POST(address string, request interface{}, output interface{}) error {

	// Prepare request.
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("could not marshal request: %w", err)
	}
	reader := bytes.NewReader(data)

	// Execute the request.
	resp, err := http.Post(address, contentTypeJSON, reader)
	if err != nil {
		return fmt.Errorf("could not execute POST request: %w", err)
	}
	defer resp.Body.Close()

	return unpackJSONResponse(resp, output)
}

// GET executes the HTTP request and unpacks the JSON response.
func GET(address string, output interface{}) error {

	resp, err := http.Get(address)
	if err != nil {
		return fmt.Errorf("could not execute GET request: %w", err)
	}
	defer resp.Body.Close()

	return unpackJSONResponse(resp, output)
}

// UnpackJSONResponse processes the HTTP response and tries to unmarshal the returned JSON into the provided struct.
func unpackJSONResponse(resp *http.Response, output interface{}) error {

	// Read the response.
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	// Check the status code.
	if resp.StatusCode == http.StatusNoContent {
		return ErrNoData
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed (code: %v, payload: %s)", resp.StatusCode, payload)
	}

	// Unpack the response data.
	err = json.Unmarshal(payload, output)
	if err != nil {
		return fmt.Errorf("could not unmarshal response (payload: %s): %w", payload, err)
	}

	return nil
}
