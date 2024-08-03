package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type bits []int

func readData(file string) []bits {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	var bData []bits
	for _, s := range sData {
		var b bits
		for _, sBit := range s {
			iBit := int(sBit - '0')
			b = append(b, iBit)
		}
		bData = append(bData, b)
	}
	return bData
}

func main() {
	data := readData("data")

	// Part One
	results := calcPower(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcLSR(data)
	fmt.Println("Part two: ", results)
}
