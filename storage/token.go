package storage

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const (
	tokenDelimiter  = ":"
	tokenFieldCount = 2
	tokenNumberBase = 10
)

// The `Storage` handler uses cursor-based pagination to iterate through the record set.
// When a batch of records is retrieved, the block number and event index of the last record
// should be provided to the `createToken` function.
// When retrieving the next batch of results, client provides the token, which is then
// decoded using `unpackToken` and used as the new reference point for
// the next batch of records.

func createToken(block string, eventIndex uint) string {
	raw := fmt.Sprintf("%v%v%v", block, tokenDelimiter, eventIndex)
	return base64.RawStdEncoding.EncodeToString([]byte(raw))
}

func unpackToken(token string) (string, uint, error) {

	// Decode base64 string.
	decoded, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return "", 0, fmt.Errorf("could not parse token: %w", err)
	}

	// Split the token fields.
	fields := strings.Split(string(decoded), tokenDelimiter)
	if len(fields) != tokenFieldCount {
		return "", 0, fmt.Errorf("unexpected token field count (have: %v, want: %v)", tokenFieldCount, len(fields))
	}

	// Verify that the first field is a valid block number.
	blockNo, ok := big.NewInt(0).SetString(fields[0], tokenNumberBase)
	if !ok {
		return "", 0, fmt.Errorf("invalid token block number (have: %v)", fields[0])
	}

	// Verify that the second field is a valid event/log index.
	eventIndex, err := strconv.ParseUint(fields[1], tokenNumberBase, 62)
	if err != nil {
		return "", 0, fmt.Errorf("invalid event index number (have: %v): %w", fields[1], err)
	}

	// Return block number as a decimal string with an event index.
	return blockNo.String(), uint(eventIndex), nil
}
