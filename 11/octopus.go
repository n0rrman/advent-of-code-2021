package main

import "math"

type octoGrid [][]octopus

type octopus struct {
	energy     int
	neighbours [](*octopus)
	flashStep  int
}

func newGrid(data [][]int) octoGrid {
	maxY := len(data)
	maxX := len(data[0])
	grid := make(octoGrid, maxY)

	for y := range data {
		grid[y] = make([]octopus, maxX)
		for x, val := range data[y] {
			var o octopus
			o.energy = val
			o.neighbours = make([](*octopus), 8)
			grid[y][x] = o
		}
	}

	for y := range grid {
		for x := range grid[y] {
			o := &grid[y][x]
			if y-1 >= 0 {
				o.neighbours[1] = &(grid[y-1][x])
				if x-1 >= 0 {
					o.neighbours[0] = &grid[y-1][x-1]
				}
				if x+1 < maxX {
					o.neighbours[2] = &grid[y-1][x+1]
				}
			}
			if y+1 < maxY {
				o.neighbours[6] = &(grid[y+1][x])
				if x-1 >= 0 {
					o.neighbours[5] = &grid[y+1][x-1]
				}
				if x+1 < maxX {
					o.neighbours[7] = &grid[y+1][x+1]
				}
			}
			if x-1 >= 0 {
				o.neighbours[3] = &grid[y][x-1]
			}
			if x+1 < maxX {
				o.neighbours[4] = &grid[y][x+1]
			}
		}
	}
	return grid
}

func (o *octopus) flash(v int) int {
	sum := 1
	o.flashStep = v
	for i := range o.neighbours {
		if o.neighbours[i] != nil {
			sum += o.neighbours[i].charge(v)
		}
	}
	return sum
}

func (o *octopus) charge(v int) int {
	if o.energy == 9 {
		o.energy = 0
		return o.flash(v)
	} else if o.flashStep != v {
		o.energy = int(math.Min(float64(o.energy+1), 9))
	}
	return 0
}
