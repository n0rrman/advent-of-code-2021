package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) []string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "")

	return sData
}

func hexToBin(data []string) []uint8 {
	var bitArray []uint8
	fourBitConv := make(map[string][]uint8)
	fourBitConv["0"] = []uint8{0, 0, 0, 0}
	fourBitConv["1"] = []uint8{0, 0, 0, 1}
	fourBitConv["2"] = []uint8{0, 0, 1, 0}
	fourBitConv["3"] = []uint8{0, 0, 1, 1}
	fourBitConv["4"] = []uint8{0, 1, 0, 0}
	fourBitConv["5"] = []uint8{0, 1, 0, 1}
	fourBitConv["6"] = []uint8{0, 1, 1, 0}
	fourBitConv["7"] = []uint8{0, 1, 1, 1}
	fourBitConv["8"] = []uint8{1, 0, 0, 0}
	fourBitConv["9"] = []uint8{1, 0, 0, 1}
	fourBitConv["A"] = []uint8{1, 0, 1, 0}
	fourBitConv["B"] = []uint8{1, 0, 1, 1}
	fourBitConv["C"] = []uint8{1, 1, 0, 0}
	fourBitConv["D"] = []uint8{1, 1, 0, 1}
	fourBitConv["E"] = []uint8{1, 1, 1, 0}
	fourBitConv["F"] = []uint8{1, 1, 1, 1}

	for _, hex := range data {
		bitArray = append(bitArray, fourBitConv[hex]...)
	}
	return bitArray
}

func main() {
	data := readData("data")
	bits := hexToBin(data)
	fmt.Println(bits)

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
