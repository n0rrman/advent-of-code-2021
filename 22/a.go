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

func rebootAndCount(i []instruction) int {
	reactor := bootUp(i)
	numOfCubes := countCubes(reactor)
	return numOfCubes
}