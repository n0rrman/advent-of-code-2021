package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	var v variables
	i := parseInstructions(data, &v)
	results := findLargestNOMAD(i, &v)

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
