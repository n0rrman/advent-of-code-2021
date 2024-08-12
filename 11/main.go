package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) [][]int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	iData := make([][]int, len(sData))
	for i := range sData {
		iData[i] = make([]int, len(sData[0]))
		for j, val := range strings.Split(sData[i], "") {
			intVal, _ := strconv.Atoi(val)
			iData[i][j] = intVal

		}
	}
	return iData
}

func main() {
	data := readData("data")

	// Part One
	results := calcFlashes(data, 100)
	fmt.Println("Part one: ", results)

	// Part Two
	results = findTotalFlash(data, 1000)
	fmt.Println("Part two: ", results)
}
