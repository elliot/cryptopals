package main

import (
	"bytes"
	"testing"
)

func TestChallenge02Output(t *testing.T) {

	input := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	result, _ := pkcs7Padding(input, 20)

	if bytes.Equal(expected, result) == false {
		t.Errorf("Set 2 - Challenge 9 Failed: Got '% x', expected '% x'", result, expected)
	}
}
