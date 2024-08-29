package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) ([]bool, [][]bool) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	values := make([]bool, 512)
	valueString := strings.Split(sData[0], "")
	for i := range values {
		values[i] = valueString[i] == "#"
	}

	trenchMap := make([][]bool, len(sData)-2)
	for y, rows := range sData[2:] {
		trenchMap[y] = make([]bool, len(rows))
		for x, val := range strings.Split(rows, "") {
			trenchMap[y][x] = val == "#"
		}
	}

	return values, trenchMap
}

func main() {
	values, trenchMap := readData("data")

	// Part One
	results := enhanceAndCalc(values, trenchMap, 2)
	fmt.Println("Part one: ", results)

	// Part Two
	results = enhanceAndCalc(values, trenchMap, 50)
	fmt.Println("Part two: ", results)
}
