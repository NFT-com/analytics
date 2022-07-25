package main

import (
	"encoding/base64"
)

func base64Encode(id string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(id))
	return encoded
}
