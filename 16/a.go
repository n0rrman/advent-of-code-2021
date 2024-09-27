package main

import (
	"fmt"
	"math"
)

type packet struct {
	version  int
	typeId   int
	lengthId int
	literal  int
}

func bitsToNum(bits []uint8) int {
	val := 0
	for i := range bits {
		val += int(bits[i]) * int(math.Pow(2, float64(len(bits)-i-1)))
	}
	return val
}

func evalLiteral(bits []uint8) int {
	i := 0
	moreBits := true
	var bitVal []uint8
	for moreBits {
		bitVal = append(bitVal, bits[i+1:i+5]...)
		moreBits = bits[i] == 1
		i += 5
	}
	return bitsToNum(bitVal)
}

func bitsToPackets(bits []uint8) packet {
	newPacket := packet{}

	index := 0
	moreBits := true
	for moreBits {
		newPacket.version = bitsToNum(bits[index : index+3])
		newPacket.typeId = bitsToNum(bits[index+3 : index+6])

		if newPacket.typeId == 4 {
			//literal, evalIndex := evalLiteral(bits[index+6:])
			//index = index + evalIndex
			newPacket.lengthId = 0
			//newPacket.literal = literal
		}
		moreBits = false
	}

	fmt.Println(newPacket.version)
	fmt.Println(newPacket.typeId)
	fmt.Println(newPacket.literal)

	return newPacket
}

// type 4: literal value => list of packs, 1 -> keep reading, 0 -> last pack
// type !4: operator => length bit = 0 => 15, 1 => next 11 decide number of bits

func recCalcVersionSum(bits []uint8) int {
	version := bitsToNum(bits[0:3])
	typeId := bitsToNum(bits[3:6])

	// type 4 -> Literal
	var literal int
	if typeId == 4 {
		literal = evalLiteral(bits[6:])
		fmt.Println("literal: ", literal)
	} else {
		lengthId := bits[6]
		var length int
		// Number of bits
		if lengthId == 0 {
			length = bitsToNum(bits[7 : 7+15])
			fmt.Println("num bits: ", length)
		}
		// Number of subpackets
		if lengthId == 1 {
			length = bitsToNum(bits[7 : 7+11])
			fmt.Println("num subpacks:", length)
		}
	}
	return 0
}

func calcVersionSum(bits []uint8) int {
	return recCalcVersionSum(bits)
}
