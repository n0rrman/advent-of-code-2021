package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcVentOverlapA(data)

	const e = 5
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcVentOverlapB(data)

	const e = 12
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
