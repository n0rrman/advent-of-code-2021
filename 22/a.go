package main

func bootUp(i []instruction) [][][]bool {
	return nil
}

func countCubes(reactor [][][]bool) int {
	count := 0
	for z := range reactor {
		for y := range reactor[z] {
			for x := range reactor[z][y] {
				if reactor[z][y][x] {
					count++
				}
			}
		}
	}
	return count
}

func normalise(i []instruction) []instruction {
	for index, instr := range i {
		for c, coord := range instr.Coords {
			if coord[0] > coord[1] {
				i[index].Coords[c] = [2]int{coord[1], coord[0]}
			}
		}
	}
	return i
}

func rebootAndCount(i []instruction) int {
	instructions := normalise(i)
	reactor := bootUp(instructions)
	numOfCubes := countCubes(reactor)
	return numOfCubes
}
