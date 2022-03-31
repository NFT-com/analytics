package storage

import (
	"encoding/base64"
	"fmt"
)

// The `Storage` handler uses cursor-based pagination to iterate through the record set.
// When a batch of records is retrieved, the ID of the last record should
// be provided to the `createToken` function.
// When retrieving the next batch of results, client provides the token, which is then
// decoded using `unpackToken` and used as the new reference point for
// the next batch of records.

func createToken(id string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(id))
}

func unpackToken(token string) (string, error) {

	decoded, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return "", fmt.Errorf("could not parse token: %w", err)
	}

	return string(decoded), nil
}
