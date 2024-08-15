package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) ([][]int, [][]int) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n\n")
	sPoints := strings.Split(sData[0], "\n")
	sFolds := strings.Split(sData[1], "\n")

	points := make([][]int, len(sPoints))
	for i := range sPoints {
		point := strings.Split(sPoints[i], ",")
		points[i] = make([]int, 2)
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		points[i][0] = x
		points[i][1] = y
	}

	folds := make([][]int, len(sFolds))
	for i := range sFolds {
		fullStr := strings.Split(sFolds[i], " ")
		dirAndNum := strings.Split(fullStr[2], "=")
		sDir := dirAndNum[0]
		num := dirAndNum[1]

		dir := 0
		if sDir == "y" {
			dir = 1
		}

		folds[i] = make([]int, 2)
		folds[i][0] = dir
		folds[i][1], _ = strconv.Atoi(num)
	}

	return points, folds
}

func main() {
	points, folds := readData("data")

	// Part One
	results := foldAndCount(points, folds, 1)
	fmt.Println("Part one: ", results)

	// Part Two
	foldAndPrint(points, folds)
}
