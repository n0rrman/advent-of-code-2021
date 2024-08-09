package main

import (
	"testing"
)

func TestA(t *testing.T) {
	_, output := readData("test_data")
	results := calc1478(output)

	const e = 26
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	input, output := readData("test_data")
	results := calcSum(input, output)

	const e = 61229
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
