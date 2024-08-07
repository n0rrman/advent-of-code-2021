package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcLanterns(data, 80)

	const e = 5934
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcLanterns(data, 256)

	const e = 26984457539
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
