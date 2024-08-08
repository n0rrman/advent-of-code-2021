package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcLoc(data)

	const e = 37
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcIncrLoc(data)

	const e = 168
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
