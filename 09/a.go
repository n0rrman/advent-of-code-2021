package main

type coords struct {
	x int
	y int
}

func findLowPoints(data [][]int) []coords {
	maxY := len(data)
	maxX := len(data[0])

	var list []coords

	for y := range data {
		for x := range data[y] {
			val := data[y][x]
			minPoint := true

			if y-1 >= 0 {
				minPoint = data[y-1][x] > val && minPoint
			}
			if x-1 >= 0 {
				minPoint = data[y][x-1] > val && minPoint
			}
			if y+1 < maxY {
				minPoint = data[y+1][x] > val && minPoint
			}
			if x+1 < maxX {
				minPoint = data[y][x+1] > val && minPoint
			}

			if minPoint {
				var c coords
				c.x = x
				c.y = y
				list = append(list, c)
			}
		}
	}
	return list
}

func calcLowPoints(data [][]int) int {
	sum := 0
	c := findLowPoints(data)
	for _, coord := range c {
		sum += data[coord.y][coord.x] + 1
	}
	return sum
}
