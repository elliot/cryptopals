package main

import (
	"testing"
)

func TestChallenge01Output(t *testing.T) {

	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result := convertHexToBase64(input)

	if expected != result {
		t.Errorf("Set 1 - Challenge 1 Failed: Got '%s', expected '%s'", result, expected)
	}
}
