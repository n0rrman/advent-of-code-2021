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

	topData := strings.Split(sData[2], "#")
	bottomData := strings.Split(sData[3], "#")

	fmt.Println(topData[3], topData[4], topData[5], topData[6])
	fmt.Println(bottomData[1], bottomData[2], bottomData[3], bottomData[4])

	rooms := [][]string{
		{topData[3], bottomData[1]},
		{topData[4], bottomData[2]},
		{topData[5], bottomData[3]},
		{topData[6], bottomData[4]},
	}
	return rooms
}

func main() {
	rooms := readData("test_data")
	fmt.Println(rooms)

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
