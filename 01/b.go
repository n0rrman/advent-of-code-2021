package main

func calcSum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func calcIncrSums(startVals []int, vals []int) int {
	prev := calcSum(startVals)
	count := 0
	for i := range vals {
		new := calcSum(vals[i : i+3])
		if new > prev {
			count += 1
		}
		prev = new
	}

	return count
}
