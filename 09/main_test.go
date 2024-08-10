package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcLowPoints(data)

	const e = 15
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcLargestBasins(data)

	const e = 1134
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
