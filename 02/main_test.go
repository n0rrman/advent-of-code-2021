package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcPos(data)

	const e = 150
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := calcPosAndAim(data)

	const e = 900
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
