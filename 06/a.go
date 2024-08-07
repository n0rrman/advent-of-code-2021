package main

func newDay(state []int) []int {
	newState := make([]int, 9)
	for i, val := range state {
		if i == 0 {
			newState[8] += val
			newState[6] += val
		} else {
			newState[i-1] += val
		}
	}

	return newState
}

func calcLanterns(initState []int, days int) int {
	state := make([]int, 9)
	for _, x := range initState {
		state[x] += 1
	}

	for i := 0; i < days; i++ {
		state = newDay(state)
	}

	count := 0
	for _, val := range state {
		count += val
	}

	return count
}
