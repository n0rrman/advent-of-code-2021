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
	test_data := readData("test_data")
	data := readData("data")

	// Part One
	test := calcIncrs(test_data[0], test_data[1:])
	results := calcIncrs(data[0], data[1:])
	if test != 7 {
		os.Exit(1)
	}
	fmt.Println("Part one: ", results)

	// Part Two
	test = calcIncrSums(test_data[0], test_data[1:])
	results = calcIncrSums(data[0], data[0:])
	fmt.Println(test)
	if test != 5 {
		os.Exit(1)
	}
	fmt.Println("Part two: ", results)
}
