package main

func traverseCave(c *cave, cs caveSystem, visited []string) int {
	visited = append(visited, c.name)

	if c.name == "end" {
		return 1
	}

	paths := 0
	for i := range c.connections {
		next := c.connections[i]
		if !in(next.name, visited) || next.bigCave {
			paths += traverseCave(next, cs, visited)
		}
	}
	return paths
}

func in(e string, list []string) bool {
	for i := range list {
		if list[i] == e {
			return true
		}
	}
	return false
}

func calcPaths(data [][]string) int {

	var visited []string
	cs := buildCaveSystem(data)

	start := cs.getCave("start")
	return traverseCave(start, cs, visited)

}
