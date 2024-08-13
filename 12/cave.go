package main

import (
	"unicode"
)

type cave struct {
	name        string
	bigCave     bool
	connections []*cave
}

type caveSystem []cave

func buildCaveSystem(data [][]string) caveSystem {
	var newCS caveSystem

	// Add all nodes
	for _, nodes := range data {
		from := nodes[0]
		to := nodes[1]

		newCS = newCS.addCave(from)
		newCS = newCS.addCave(to)
	}

	// Add all edges
	for _, edge := range data {
		from := edge[0]
		to := edge[1]

		fromPointer := newCS.getCave(from)
		toPointer := newCS.getCave(to)

		toPointer.connections = append(toPointer.connections, fromPointer)
		fromPointer.connections = append(fromPointer.connections, toPointer)
	}

	return newCS
}

func (c caveSystem) addCave(name string) caveSystem {
	for i := range c {
		if c[i].name == name {
			return c
		}
	}

	var newCave cave
	newCave.name = name
	if unicode.IsUpper(rune(name[0])) {
		newCave.bigCave = true
	} else {
		newCave.bigCave = false
	}
	return append(c, newCave)
}

func (c caveSystem) getCave(name string) *cave {
	for i := range c {
		if c[i].name == name {
			return &c[i]
		}
	}
	return nil
}
