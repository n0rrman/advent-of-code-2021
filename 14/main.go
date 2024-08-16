package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) (string, [][]string) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(strings.Replace(string(body), "\n\n", "\n", -1)), "\n")
	str := sData[0]
	pairs := make([][]string, len(sData)-1)

	for i := 1; i <= len(pairs); i++ {
		pairs[i-1] = make([]string, 2)
		pair := strings.Split(sData[i], " -> ")
		pairs[i-1][0], pairs[i-1][1] = pair[0], pair[1]
	}

	return str, pairs
}

func main() {
	str, pairs := readData("data")

	// Part One
	results := calcPolyValue(str, pairs)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 2
	fmt.Println("Part two: ", results)
}
