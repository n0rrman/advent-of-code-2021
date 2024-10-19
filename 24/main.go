package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) []string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	return sData
}

type instructionType int

const (
	Inp instructionType = iota + 1
	Add
	Mul
	Div
	Mod
	Eql
)

type instruction struct {
	instr instructionType
	vars  []int
}

func parseInstructions(ss []string) []instruction {
	var i []instruction
	for _, s := range ss {
		split := strings.Split(s, " ")
		switch split[0] {
		case "inp":
			fmt.Println("Inp")
		case "add":
			fmt.Println("Add")
		case "mul":
			fmt.Println("Mul")
		case "div":
			fmt.Println("Div")
		case "mod":
			fmt.Println("Mod")
		case "eql":
			fmt.Println("Eql")
		}

		i = append(i, instruction{})
	}
	// inp a
	// add a b
	// mul a b
	// div a b
	// mod a b
	// eql a b
	return i
}

func main() {
	data := readData("test_data")

	// Part One
	i := parseInstructions(data)
	results := findLargestNOMAD(i)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 1
	fmt.Println("Part two: ", results)
}
