package main

import (
	"testing"
)

func TestA(t *testing.T) {
	points, folds := readData("test_data")
	results := foldAndCount(points, folds, 1)

	const e = 17
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	points, folds := readData("test_data")
	results := foldAndCount(points, folds, 1)

	const e = 17
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
