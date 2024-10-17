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
}

func parseInstructions([]string) []instruction {
	// inp a
	// add a b
	// mul a b
	// div a b
	// mod a b
	// eql a b
	return []instruction{}
}

func main() {
	data := readData("test_data")

	// Part One
	results := findLargestNOMAD(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 1
	fmt.Println("Part two: ", results)
}
