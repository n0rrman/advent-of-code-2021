package main

import (
	"testing"
)

func TestA(t *testing.T) {
	str, pairs := readData("test_data")
	results := calcPolyValue(str, pairs, 10)

	const e = 1588
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	str, pairs := readData("test_data")
	results := calcPolyValue(str, pairs, 40)

	const e = 2188189693529
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
