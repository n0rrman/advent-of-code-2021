package main

import (
	"testing"
)

func TestA(t *testing.T) {
	randomNumData, boardsData := readData("test_data")
	results := calcFirstWin(randomNumData, boardsData)

	const e = 1
	if results != 4512 {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	randomNumData, boardsData := readData("test_data")
	results := calcLastWin(randomNumData, boardsData)

	const e = 1924
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
