package main

import (
	"encoding/hex"
)

func fixedXOR(input, key string) string {

	inputBytes, err := hex.DecodeString(input)
	if err != nil {
		return ""
	}

	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return ""
	}

	if len(inputBytes) != len(keyBytes) {
		return ""
	}

	length := len(inputBytes)

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = inputBytes[i] ^ keyBytes[i]
	}

	return hex.EncodeToString(result)
}
