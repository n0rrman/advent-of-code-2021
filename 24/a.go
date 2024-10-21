package main

import "fmt"

func executeProgram(instrs *[]instruction, n []int) {
	counter := 0
	for i, instr := range *instrs {
		switch instr.instr {
		case Inp:
			*(*instrs)[i].varA = n[counter]
			counter++
		case Add:
			*(*instrs)[i].varA = *instr.varA + *instr.varB
		case Mul:
			*(*instrs)[i].varA = *instr.varA * *instr.varB
		case Div:
			*(*instrs)[i].varA = *instr.varA / *instr.varB
		case Mod:
			*(*instrs)[i].varA = *instr.varA % *instr.varB
		case Eql:
			*(*instrs)[i].varA = *instr.varA / *instr.varB
			if *instr.varA == *instr.varB {
				*(*instrs)[i].varA = 1
			} else {
				*(*instrs)[i].varA = 0
			}
		}
	}
}

func clearCache(v *variables) {
	v.w = 0
	v.x = 0
	v.y = 0
	v.z = 0
}

func findLargestNOMAD(i *[]instruction, v *variables) int {

	for idx := 1; idx < 15; idx++ {
		clearCache(v)
		numbers := []int{idx}
		executeProgram(i, numbers)
		fmt.Println(*v)
	}

	return 0
}

// instruction type
//		instruct enum -> 2 params
// build program
// 14 inp instrcutions -> 14 digits
