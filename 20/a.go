package main

import (
	"fmt"
	"math"
)

func calcLit(trenchMap [][]bool) int {
	count := 0
	for _, rows := range trenchMap {
		for _, val := range rows {
			if val {
				count++
			}
		}
	}
	return count
}

func enhanceMap(values []bool, extendedMap [][]bool) [][]bool {
	newMap := make([][]bool, len(extendedMap)-2)
	for i := range newMap {
		newMap[i] = make([]bool, len(extendedMap[0])-2)
	}

	for y := 1; y < len(extendedMap)-1; y++ {
		for x := 1; x < len(extendedMap[y])-1; x++ {
			val := isLit([]bool{
				extendedMap[y-1][x-1], extendedMap[y-1][x], extendedMap[y-1][x+1],
				extendedMap[y][x-1], extendedMap[y][x], extendedMap[y][x+1],
				extendedMap[y+1][x-1], extendedMap[y+1][x], extendedMap[y+1][x+1],
			})
			newMap[y-1][x-1] = values[val]
		}
	}
	return newMap
}

func isLit(bits []bool) int {
	sum := 0
	for i, val := range bits {
		if val {
			x := int(math.Pow(2, float64(len(bits)-i-1)))
			sum += x
		}
	}
	return sum
}

func extendMap(oldMap [][]bool, void bool) [][]bool {
	extendedMap := make([][]bool, len(oldMap)+6)
	for i := range extendedMap {
		extendedMap[i] = make([]bool, len(oldMap[0])+6)
		for j := range extendedMap[i] {
			extendedMap[i][j] = void
			if j >= 3 && j < len(extendedMap)-3 && i >= 3 && i < len(extendedMap[0])-3 {
				extendedMap[i][j] = oldMap[i-3][j-3]
			}
		}
	}
	return extendedMap
}

func printMap(printMap [][]bool) {
	for y := range printMap {
		for x := range printMap[y] {
			if printMap[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func enhanceAndCalc(values []bool, trenchMap [][]bool, iter int) int {
	var extendedMap [][]bool
	enhancedMap := trenchMap
	void := false

	for i := 0; i < iter; i++ {
		extendedMap = extendMap(enhancedMap, void)
		enhancedMap = enhanceMap(values, extendedMap)
		void = enhancedMap[0][0]
	}

	return calcLit(enhancedMap)
}
