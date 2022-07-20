package api

import (
	"encoding/base64"
	"fmt"
)

// createCursor Base64 encodes the provided ID, typically used for the pagination cursor.
func createCursor(id string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(id))
	return encoded
}

// decodeCursor decodes the Base64-encoded ID, typically provided as the pagination cursor.
func decodeCursor(cursor string) (string, error) {

	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", fmt.Errorf("could not decode cursor: %w", err)
	}

	return string(decoded), nil
}
