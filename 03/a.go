package main

import (
	"math"
)

func calcPower(b []bits) int {

	gamma := 0
	epsilon := 0
	bitLen := len(b[0])

	for i := range b[0] {
		count := 0
		for _, bits := range b {
			count += bits[i]
		}
		isGamma := count > (len(b) / 2)

		bitVal := int(math.Pow(2, float64((bitLen - 1 - i))))

		if isGamma {
			gamma += bitVal
		} else {
			epsilon += bitVal
		}
	}

	return gamma * epsilon
}
