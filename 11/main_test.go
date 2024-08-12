package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcFlashes(data, 100)

	const e = 1656
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := findTotalFlash(data, 1000)

	const e = 195
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
