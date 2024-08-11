package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcSysErr(data)

	const e = 26397
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcAutocompScore(data)

	const e = 288957
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
