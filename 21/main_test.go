package main

import (
	"testing"
)

func TestDice(t *testing.T) {
	var d dice = 5
	d.rollDice()
	d.rollDice()

	if d.readDice() != 7 {
		t.Errorf("Error rolling dice")
	}
}

func TestDataLoading(t *testing.T) {
	p1, p2 := readData("test_data")

	if p1 != 4 && p2 != 8 {
		t.Errorf("Error loading the test data")
	}

	p1, p2 = readData("data")

	if p1 != 10 && p2 != 2 {
		t.Errorf("Error loading the data")
	}
}

func TestA(t *testing.T) {
	p1, p2 := readData("test_data")
	results := playAndCount(p1, p2)

	const e = 739785
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
