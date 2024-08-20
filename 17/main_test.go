package main

import (
	"testing"
)

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
