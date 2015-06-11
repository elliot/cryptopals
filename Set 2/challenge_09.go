package main

import (
	"bytes"
	"fmt"
)

func pkcs7Padding(data []byte, blockSize int) ([]byte, error) {

	if blockSize <= 0 {
		return nil, fmt.Errorf("Invalid block size %d", blockSize)
	}

	// Determine the length to pad
	padLength := 1

	for ((len(data) + padLength) % blockSize) != 0 {
		padLength++
	}

	// Create a slice of bytes padded with the length
	padding := bytes.Repeat([]byte{byte(padLength)}, padLength)

	return append(data, padding...), nil
}
