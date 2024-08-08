package main

import "math"

func calcIncrLoc(values []int) int {
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
		prev := 0
		count := 0
		for _, x := range values {
			val := int(math.Abs(float64(x - i)))
			count += val * (val + 1) / 2
		}
		if prev > count {
			break
		}
		prev = count
		results[i] = count
	}

	minimum := results[0]
	for _, val := range results {
		if val < minimum {
			minimum = val
		}
	}
	return minimum
}
