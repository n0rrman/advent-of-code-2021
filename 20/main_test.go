package main

import (
	"testing"
)

func TestValues(t *testing.T) {
	values, _ := readData("test_data")

	res := make([]bool, 3)
	res[0] = values[34]
	res[1] = values[35]
	res[2] = values[36]

	if !res[0] {
		t.Errorf("Expected true!")
	}
	if !res[1] {
		t.Errorf("Expected true!")
	}
	if res[2] {
		t.Errorf("Expected false!")
	}
}

func TestIsLit(t *testing.T) {
	res := isLit([]bool{false, false, false, true, false, false, false, true, false})

	const e = 34
	if res != e {
		t.Errorf("Expected %v, but got %v", e, res)
	}
}

func TestA(t *testing.T) {
	values, trenchMap := readData("test_data")

	results := enhanceAndCalc(values, trenchMap, 2)

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
