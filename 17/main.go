package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) ([]int, []int) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), " ")
	xString := strings.Split(strings.Split(sData[2], "=")[1], "..")
	yString := strings.Split(strings.Split(sData[3], "=")[1], "..")

	x, y := make([]int, 2), make([]int, 2)
	x[0], _ = strconv.Atoi(xString[0])
	x[1], _ = strconv.Atoi(xString[1])
	y[0], _ = strconv.Atoi(yString[0])
	y[1], _ = strconv.Atoi(yString[1])

	return x, y
}

func main() {
	_, _ = readData("data")

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
