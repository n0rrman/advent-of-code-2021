package main

func calcIncrs(startVal int, vals []int) int {
	prev := startVal
	count := 0
	for _, d := range vals {
		if d > prev {
			count += 1
		}
		prev = d
	}

	return count
}
