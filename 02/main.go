package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	method string
	count  int
}

func readData(file string) []instruction {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	var iData []instruction
	for _, d := range sData {
		split := strings.Split(d, " ")
		method := split[0]
		count, _ := strconv.Atoi(split[1])
		iData = append(iData, instruction{method, count})
	}
	return iData
}

func main() {
	data := readData("data")

	// Part One
	results := calcPos(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = calcPosAndAim(data)
	fmt.Println("Part two: ", results)
}
