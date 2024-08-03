package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcPower(data)

	const e = 198
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcLSR(data)

	const e = 230
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
