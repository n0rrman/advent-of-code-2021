package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcPaths(data)

	const e = 19
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := recalcPaths(data)

	const e = 103
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
