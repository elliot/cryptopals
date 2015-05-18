package main

import (
	"encoding/base64"
	"encoding/hex"
)

// Set 1 / Challenge 1 / Convert hex to base64

func convertHexToBase64(input string) string {

	bytes, err := hex.DecodeString(input)

	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytes)
}
