package main

import "fmt"

type coords struct {
	x int
	y int
}

type workList map[coords]int

func makeVisitedTable(data [][]int) [][]bool {
	visited := make([][]bool, len(data))
	for y, row := range data {
		visited[y] = make([]bool, len(row))
		for x := range data {
			visited[y][x] = false
		}
	}
	return visited
}

func addVisited(visited [][]bool, c coords) [][]bool {
	visited[c.y][c.x] = true
	return visited
}

func isValid(visited [][]bool, c coords) bool {
	if (c.y < 0) || (c.y >= len(visited)) {
		return false
	}
	if (c.x < 0) || (c.x > len(visited[0])) {
		return false
	}
	return !visited[c.y][c.x]
}

func addNeighboursToList(wl workList, c coords, energy int, data [][]int, visited [][]bool) workList {
	e := coords{x: c.x + 1, y: c.y}
	w := coords{x: c.x - 1, y: c.y}
	n := coords{x: c.x, y: c.y + 1}
	s := coords{x: c.x, y: c.y - 1}

	if isValid(visited, n) {
		wl[n] = energy + data[n.y][n.x]
	}
	if isValid(visited, s) {
		wl[s] = energy + data[s.y][s.x]
	}
	if isValid(visited, e) {
		wl[e] = energy + data[e.y][e.x]
	}
	if isValid(visited, w) {
		wl[w] = energy + data[w.y][w.x]
	}

	fmt.Println(wl)

	return wl
}

func isEmpty(wl workList) bool {
	return len(wl) <= 0
}

func eq(a, b coords) bool {
	return a.x == b.x && a.y == b.y
}

func calcLowestRisk(data [][]int) int {
	wl := make(workList)
	energySpent := 0
	visited := makeVisitedTable(data)
	end := coords{x: len(data[0]), y: len(data)}

	visited = addVisited(visited, coords{0, 0})
	wl = addNeighboursToList(wl, coords{0, 0}, energySpent, data, visited)

	fmt.Println(isEmpty(wl), wl)

	for !isEmpty(wl) {
		for k, v := range wl {
			if v == energySpent {
				delete(wl, k)
				visited = addVisited(visited, k)
				wl = addNeighboursToList(wl, k, energySpent, data, visited)
			}
			if eq(k, end) {
				return energySpent
			}
		}
		energySpent++
	}

	fmt.Println("Ooops, something went wrong...")
	return 0
}
