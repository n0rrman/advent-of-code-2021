package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestHexToBin(t *testing.T) {
	hex := strings.Split("38006F45291200", "")
	converted := hexToBin(hex)

	bin := strings.Split("00111000000000000110111101000101001010010001001000000000", "")

	fullMatch := true
	for i := range bin {
		bit, _ := strconv.Atoi(bin[i])
		if converted[i] != uint8(bit) {
			fullMatch = false
			break
		}
	}

	if !fullMatch {
		t.Error("Elements does not match!")
	}
}

func TestA(t *testing.T) {
	data := readData("test_data")
	_ = data
	results := 0

	const e = 1
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	results := 0

	const e = 2
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
