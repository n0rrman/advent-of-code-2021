package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) [][][]int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n\n")
	scanners := make([][][]int, len(sData))

	for k, scannerDataString := range sData {
		scannerData := strings.Split(scannerDataString, "\n")
		scanners[k] = make([][]int, len(scannerData)-1)
		for i, row := range scannerData {
			if i == 0 {
				continue
			}
			intRow := make([]int, 3)
			for j, val := range strings.Split(row, ",") {
				intRow[j], _ = strconv.Atoi(val)
			}
			scanners[k][i-1] = intRow
		}
	}
	return scanners
}

func main() {
	data := readData("test_data")

	for _, scanner := range data {
		fmt.Println(scanner)

	}

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
