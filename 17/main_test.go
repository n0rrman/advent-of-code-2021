package main

import (
	"testing"
)

func TestReadData(t *testing.T) {
	x, y := readData("test_data")

	if x[0] != 20 && x[1] != 30 {
		t.Error("Error loading x values")
	}
	if y[0] != -10 && y[1] != -5 {
		t.Error("Error loading y values")
	}
}

func TestA(t *testing.T) {
	results := 0

	const e = 45
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
