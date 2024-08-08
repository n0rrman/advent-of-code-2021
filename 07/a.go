package main

import (
	"math"
)

func calcLoc(values []int) int {

	min, max := values[0], values[0]
	for _, val := range values {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}

	results := make([]int, max+1)
	for i := min; i <= max; i++ {
		count := 0
		prev := 0
		for _, x := range values {
			count += int(math.Abs(float64(x - i)))
		}
		if prev > count {
			break
		}
		prev = count
		results[i] = count
	}

	minimum := results[0]
	for _, x := range results {
		if x < minimum {
			minimum = x
		}
	}
	return minimum

}
