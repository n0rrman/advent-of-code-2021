package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) [][]int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	data := make([][]int, len(sData))

	for y := range data {
		data[y] = make([]int, len(sData[y]))
		for x, val := range strings.Split(sData[y], "") {
			var code int
			switch val {
			case ".":
				code = 0
			case ">":
				code = 1
			case "v":
				code = 2
			}
			data[y][x] = code
		}
	}
	return data
}

func main() {
	data := readData("test_data")
	fmt.Println(data)

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
