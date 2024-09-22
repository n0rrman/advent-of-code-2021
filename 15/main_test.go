package main

import (
	"testing"
)

func TestMakeVisitedTable(t *testing.T) {
	data := readData("test_data")
	visited := makeVisitedTable(data)

	if len(visited) != len(data) {
		t.Errorf("Y length mismatch")
	}
	if len(visited[0]) != len(data[0]) {
		t.Errorf("X length mismatch")
	}
	if visited[1][5] {
		t.Errorf("Value should be false")
	}
}

func TestEq(t *testing.T) {
	a := coords{1, 2}
	b := coords{2, 1}
	c := coords{1, 2}

	if eq(a, b) {
		t.Errorf("a & b should not be equal")
	}
	if !eq(a, c) {
		t.Errorf("a & c should be equal")
	}
	if eq(b, c) {
		t.Errorf("b & c should not be equal")
	}

}

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcLowestRisk(data)

	const e = 40
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("data")
	results := calcLowestRiskExtended(data)

	const e = 315
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
