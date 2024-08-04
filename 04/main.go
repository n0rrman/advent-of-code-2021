package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) ([]int, [][][]int) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n\n")

	var randomNums []int
	for _, n := range strings.Split(sData[0], ",") {
		i, _ := strconv.Atoi(n)
		randomNums = append(randomNums, i)
	}

	var boards [][][]int
	for _, data := range sData[1:] {
		var newCol [][]int
		for _, row := range strings.Split(data, "\n") {
			var newRow []int
			for _, val := range strings.Split(strings.Replace(strings.Trim(row, " "), "  ", " ", -1), " ") {
				i, _ := strconv.Atoi(val)
				newRow = append(newRow, i)
			}
			newCol = append(newCol, newRow)
		}
		boards = append(boards, newCol)
	}

	return randomNums, boards
}

func main() {
	randomNumData, boardsData := readData("data")

	// Part One
	results := calcFirstWin(randomNumData, boardsData)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcLastWin(randomNumData, boardsData)
	fmt.Println("Part two: ", results)
}
