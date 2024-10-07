package main

import (
	"testing"
)

func TestA1(t *testing.T) {
	data := readData("test_data1")
	results := rebootAndCount(data)

	const e = 39
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestA2(t *testing.T) {
	data := readData("test_data2")
	results := rebootAndCount(data)

	const e = 590784
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
