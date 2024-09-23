package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) [][]int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	data := make([][]int, len(sData))
	for y := range sData {
		data[y] = make([]int, len(sData[y]))
		for x := range strings.Split(sData[y], "") {
			data[y][x] = int(sData[y][x]) - 48
		}
	}
	return data
}

func main() {
	data := readData("data")

	// Part One
	results := calcLowestRisk(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcLowestRiskExtended(data)
	fmt.Println("Part two: ", results)
}
