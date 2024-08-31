package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) (int, int) {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	playerOne, _ := strconv.Atoi(strings.Split(sData[0], " ")[4])
	playerTwo, _ := strconv.Atoi(strings.Split(sData[1], " ")[4])

	return playerOne, playerTwo
}

func main() {
	p1, p2 := readData("data")
	fmt.Println(p1, p2)

	// Part One
	results := playAndCount(p1, p2)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 2
	fmt.Println("Part two: ", results)
}
