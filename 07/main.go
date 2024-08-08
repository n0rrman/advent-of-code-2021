package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) []int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), ",")

	iData := make([]int, len(sData))
	for i := range sData {
		iData[i], _ = strconv.Atoi(sData[i])
	}

	return iData
}

func main() {
	data := readData("data")

	// Part One
	results := calcLoc(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcIncrLoc(data)
	fmt.Println("Part two: ", results)
}
