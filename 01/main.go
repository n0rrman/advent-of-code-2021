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
	sData := strings.Split(string(body[:]), "\n")
	var iData []int
	for _, d := range sData {
		i, _ := strconv.Atoi(d)
		iData = append(iData, i)
	}
	return iData
}

func main() {
	data := readData("data")

	// Part One
	results := calcIncrs(data[0], data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcIncrSums(data[:4], data)
	fmt.Println("Part two: ", results)
}
