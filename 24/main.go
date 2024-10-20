package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) []string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	return sData
}

func parseInstructions(ss []string, v *variables) []instruction {
	var i []instruction

	for _, s := range ss {
		// Instruction
		split := strings.Split(s, " ")
		var newInstr instruction
		switch split[0] {
		case "inp":
			newInstr.instr = Inp
		case "add":
			newInstr.instr = Add
		case "mul":
			newInstr.instr = Mul
		case "div":
			newInstr.instr = Div
		case "mod":
			newInstr.instr = Mod
		case "eql":
			newInstr.instr = Eql
		}

		// First variable
		val, err := strconv.Atoi(split[1])
		if err == nil {
			newInstr.constA = val
			newInstr.varA = &newInstr.constA
		} else {
			switch split[1] {
			case "w":
				newInstr.varA = &v.w
			case "x":
				newInstr.varA = &v.x
			case "y":
				newInstr.varA = &v.y
			case "z":
				newInstr.varA = &v.z
			}
		}

		// Second variable
		if len(split) > 2 {
			val, err := strconv.Atoi(split[2])
			if err == nil {
				newInstr.constB = val
				newInstr.varB = &newInstr.constB
			} else {
				switch split[2] {
				case "w":
					newInstr.varB = &v.w
				case "x":
					newInstr.varB = &v.x
				case "y":
					newInstr.varB = &v.y
				case "z":
					newInstr.varB = &v.z
				}
			}
		}

		i = append(i, newInstr)
	}
	return i
}

func main() {
	data := readData("test_data")
	var v variables
	i := parseInstructions(data, &v)

	// Part One
	results := findLargestNOMAD(i, &v)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 1
	fmt.Println("Part two: ", results)
}
