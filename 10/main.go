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

	rData := make([][]string, len(sData))
	for i := range sData {
		rData[i] = make([]string, len(sData[i]))
		for j, val := range strings.Split(sData[i], "") {
			rData[i][j] = val
		}
	}
	return rData

}

func main() {
	data := readData("data")

	// Part One
	results := calcSysErr(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcAutocompScore(data)
	fmt.Println("Part two: ", results)
}
