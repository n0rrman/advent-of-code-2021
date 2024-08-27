package main

import (
	"testing"
)

func TestA(t *testing.T) {
	values, trenchMap := readData("data")

	results := 0
	_ = values
	_ = trenchMap

	const e = 35
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
