package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordPair struct {
	fromX int
	fromY int
	toX   int
	toY   int
}

func readData(file string) []coordPair {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	var cData []coordPair
	for _, row := range sData {
		data := strings.Split(row, " -> ")
		from := strings.Split(data[0], ",")
		to := strings.Split(data[1], ",")

		var pair coordPair
		pair.fromX, _ = strconv.Atoi(from[0])
		pair.fromY, _ = strconv.Atoi(from[1])
		pair.toX, _ = strconv.Atoi(to[0])
		pair.toY, _ = strconv.Atoi(to[1])
		cData = append(cData, pair)
	}

	return cData
}

func main() {
	data := readData("data")

	// Part One
	results := calcVentOverlapA(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcVentOverlapA(data)
	fmt.Println("Part two: ", results)
}
