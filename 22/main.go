package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type Range [2]int 
type instruction struct {
	Mode bool
	Coords [3]Range
}


func readData(file string) []instruction {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	instructions := make([]instruction, len(sData))
	for i, row := range sData {
		modeSplit := strings.Split(row, " ")
		
		instructions[i] = instruction{}	
		instructions[i].Mode = modeSplit[0] == "on"

		for c, coord := range strings.Split(modeSplit[1], ",") {
			cleanUp := strings.Split(coord, "=")
			coords := strings.Split(cleanUp[1], "..")
			c1, _ := strconv.Atoi(coords[0])
			c2, _ := strconv.Atoi(coords[1])
			instructions[i].Coords[c] = Range{c1, c2}
		
		}
	}

	return instructions
}

func main() {
	data := readData("test_data2")
	fmt.Println(data)
	_ = data

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
