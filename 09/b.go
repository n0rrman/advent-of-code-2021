package main

import (
	"sort"
)

func calcLargestBasins(data [][]int) int {
	c := findLowPoints(data)
	maxVals := make([]int, 3)
	for i := range maxVals {
		maxVals[i] = 0
	}

	for _, point := range c {
		var queue []coords
		visited := make([][]bool, len(data))
		for i := range visited {
			visited[i] = make([]bool, len(data[0]))
		}

		count := 0
		current := point
		run := true
		for run {
			y := current.y
			x := current.x
			visited[y][x] = true

			if y-1 >= 0 {
				if data[y-1][x] != 9 && !visited[y-1][x] {
					var c coords
					c.x = x
					c.y = y - 1
					queue = append(queue, c)
				}
			}
			if y+1 < len(data) {
				if data[y+1][x] != 9 && !visited[y+1][x] {
					var c coords
					c.x = x
					c.y = y + 1
					queue = append(queue, c)
				}
			}
			if x+1 < len(data[0]) {
				if data[y][x+1] != 9 && !visited[y][x+1] {
					var c coords
					c.x = x + 1
					c.y = y
					queue = append(queue, c)
				}
			}
			if x-1 >= 0 {
				if data[y][x-1] != 9 && !visited[y][x-1] {
					var c coords
					c.x = x - 1
					c.y = y
					queue = append(queue, c)
				}
			}

			if len(queue) > 0 {
				current = queue[0]
				queue = queue[1:]
			} else {
				run = false

				for i := range visited {
					for _, val := range visited[i] {
						if val {
							count++
						}
					}
				}

			}
		}
		if count > maxVals[0] {
			maxVals[0] = count
			sort.Ints(maxVals)
		}
	}
	return maxVals[2] * maxVals[1] * maxVals[0]
}
