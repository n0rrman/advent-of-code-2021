package main

import (
	"fmt"
	"math"
)

type packet struct {
	version int
	typeId  int
	length  int
}

func bitsToNum(bits []uint8) int {
	val := 0
	for i := range bits {
		val += int(bits[i]) * int(math.Pow(2, float64(len(bits)-i-1)))
	}
	return val
}

func bitsToPackets(bits []uint8) packet {
	newPacket := packet{}

	startIndex := 0
	moreBits := true
	for moreBits {
		newPacket.version = bitsToNum(bits[startIndex : startIndex+3])
		newPacket.typeId = bitsToNum(bits[startIndex+3 : startIndex+6])
		//for _, bit := range bits {

		fmt.Println(bits[startIndex : startIndex+3])
		fmt.Println(bits[startIndex+3 : startIndex+6])
		fmt.Println(bits[startIndex : startIndex+6])
		//packetIndex++
		moreBits = false
	}

	fmt.Println(newPacket.version)
	fmt.Println(newPacket.typeId)

	return newPacket
}

// type 4: literal value => list of packs, 1 -> keep reading, 0 -> last pack
// type !4: operator => length bit = 0 => 15, 1 => next 11 decide number of bits

func calcVersionSum(bits []uint8) int {

	packets := bitsToPackets(bits)
	_ = packets

	return 0
}
