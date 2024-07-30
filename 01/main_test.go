package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcIncrs(data[0], data)

	if results != 7 {
		t.Errorf("Expected 7, but got %v", results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcIncrSums(data[:4], data)

	if results != 5 {
		t.Errorf("Expected 5, but got %v", results)
	}
}
