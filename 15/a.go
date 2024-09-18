package main

/*
func findSafestPath(x, y int, data [][]int, visited [][]bool) int {

	if y == len(data)-1 && x == len(data[y])-1 {
		return data[y][x]
	}

	n := y > 0
	e := x < len(data[0])-1
	s := y < len(data)-1
	w := x > 0
	visited[y][x] = true
	pseudoMaxVal := 99999999999
	minVal := pseudoMaxVal

	if n && !visited[y-1][x] {
		val := findSafestPath(x, y-1, data, visited)
		minVal = int(math.Min(float64(minVal), float64(val)))
	}
	if e && !visited[y][x+1] {
		val := findSafestPath(x+1, y, data, visited)
		minVal = int(math.Min(float64(minVal), float64(val)))
	}
	if s && !visited[y+1][x] {
		val := findSafestPath(x, y+1, data, visited)
		minVal = int(math.Min(float64(minVal), float64(val)))
	}
	if w && !visited[y][x-1] {
		val := findSafestPath(x-1, y, data, visited)
		minVal = int(math.Min(float64(minVal), float64(val)))
	}
	return data[y][x] + minVal
}
*/

func calcLowestRisk(data [][]int) int {
	// Add starting node visited
	// Add all neighbours to queue
	// Energy spent = 0

	// if queue value = 0
	//		add to visited, add neighbours to queue

	// increment energy spent
	// decrement all values in queue

	return 0
}
