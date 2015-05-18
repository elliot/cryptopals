package main

import (
	"testing"
)

func TestChallenge02Output(t *testing.T) {

	input := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	result := fixedXOR(input, key)

	if expected != result {
		t.Errorf("Set 1 - Challenge 2 Failed: Got '%s', expected '%s'", result, expected)
	}
}
