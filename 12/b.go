package main

func reTraverseCave(c *cave, cs caveSystem, visited []string, dblVisit bool) int {
	visited = append(visited, c.name)

	if c.name == "end" {
		return 1
	}

	paths := 0
	for i := range c.connections {
		next := c.connections[i]
		if next.bigCave || !in(next.name, visited) {
			paths += reTraverseCave(next, cs, visited, dblVisit)
		} else if !dblVisit && next.name != "start" {
			paths += reTraverseCave(next, cs, visited, true)
		}
	}
	return paths
}

func recalcPaths(data [][]string) int {

	var visited []string
	cs := buildCaveSystem(data)

	start := cs.getCave("start")
	return reTraverseCave(start, cs, visited, false)

}
