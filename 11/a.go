package main

func calcFlashes(data [][]int, steps int) int {

	grid := newGrid(data)
	sum := 0

	for i := 1; i <= steps; i++ {
		for y := range grid {
			for x := range grid[y] {
				sum += grid[y][x].charge(i)
			}
		}
	}

	return sum
}
