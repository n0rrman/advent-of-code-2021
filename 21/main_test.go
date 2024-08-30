package main

import (
	"testing"
)

func TestDataLoading(t *testing.T) {
	p1, p2 := readData("test_data")

	if p1 != 4 && p2 != 8 {
		t.Errorf("Error loading the test data")
	}

	p1, p2 = readData("data")

	if p1 != 10 && p2 != 2 {
		t.Errorf("Error loading the data")
	}
}

func TestA(t *testing.T) {
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
