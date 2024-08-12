package main

func findTotalFlash(data [][]int, maxSteps int) int {

	grid := newGrid(data)
	size := len(data) * len(data[0])

	for i := 1; i <= maxSteps; i++ {
		sum := 0
		for y := range grid {
			for x := range grid[y] {
				sum += grid[y][x].charge(i)
			}
		}
		if sum == size {
			return i
		}
	}

	return -1
}
