package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type digit [7]bool

func readData(file string) ([][]digit, [][]digit) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	input := make([][]digit, len(sData))
	output := make([][]digit, len(sData))

	for i, val := range sData {

		row := strings.Split(val, " | ")
		sInput := row[0]
		sOutput := row[1]

		letterToNum := map[rune]int{
			'a': 0,
			'b': 1,
			'c': 2,
			'd': 3,
			'e': 4,
			'f': 5,
			'g': 6,
		}

		sInputList := strings.Split(sInput, " ")
		input[i] = make([]digit, len(sInputList))
		for j, chars := range sInputList {
			for _, char := range chars {
				input[i][j][letterToNum[char]] = true
			}
		}

		sOutputList := strings.Split(sOutput, " ")
		output[i] = make([]digit, len(sOutputList))
		for j, chars := range sOutputList {
			for _, char := range chars {
				output[i][j][letterToNum[char]] = true
			}
		}
	}

	return input, output
}

func main() {
	input, output := readData("data")

	// Part One
	results := calc1478(output)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcSum(input, output)
	fmt.Println("Part two: ", results)
}
