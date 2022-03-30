package api

import (
	"encoding/base64"
	"fmt"
)

// The API is using cursor-based pagination to iterate through the record set.
// When a batch of records is retrieved, the timestamp of the last record should
// be provided to the `createPaginationToken` function. The function will return
// the token for the next page (base64-encoding of the timestamp).
// When requesting the next page of results, the client should provide the token,
// which is then decoded using `unpackPaginationToken` and used as the new
// reference point for the next batch of records.

// FIXME: These will use time.Time in the future.
// FIXME: Rethink the names and all of this.

func createPaginationToken(timestamp string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(timestamp))
}

func unpackPaginationToken(token string) (string, error) {

	decoded, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return "", fmt.Errorf("could not parse paging token: %w", err)
	}

	return string(decoded), nil
}
