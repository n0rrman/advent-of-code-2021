package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) [][]string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	paths := make([][]string, len(sData))
	for i := range sData {
		paths[i] = strings.Split(sData[i], "-")
	}

	return paths
}

func main() {
	data := readData("data")

	// Part One
	results := calcPaths(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = recalcPaths(data)
	fmt.Println("Part two: ", results)
}
