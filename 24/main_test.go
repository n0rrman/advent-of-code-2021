package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	i := parseInstructions(data)
	results := findLargestNOMAD(i)

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
