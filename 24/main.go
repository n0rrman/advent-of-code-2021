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

type instruction struct {
	instr int
	vars  []int
}

func parseInstructions(ss []string) []instruction {
	var i []instruction
	for _, s := range ss {
		split := strings.Split(s, " ")
		fmt.Println(split[0])
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
