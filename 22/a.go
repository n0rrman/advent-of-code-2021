package main

import "fmt"

func bootUp(i []instruction, size []int) [][][]bool {
	fmt.Println("Empty")
	// Empty reactor
	reactor := make([][][]bool, size[0]+1)
	for z := range reactor {
		reactor[z] = make([][]bool, size[1]+1)
		for y := range reactor[z] {
			fmt.Println(y, size[2])
			reactor[z][y] = make([]bool, size[2]+1)
			//for x := range reactor[z][y] {
			//	reactor[z][y][x] = false
			//}
		}
	}

	fmt.Println("Fill")
	// Fill reactor
	for _, instruction := range i {
		coords := instruction.Coords
		for z := coords[0][0]; z <= coords[0][1]; z++ {
			for y := coords[1][0]; y <= coords[1][1]; y++ {
				for x := coords[2][0]; x <= coords[2][1]; x++ {
					reactor[z][y][x] = instruction.Mode
				}
			}
		}
	}

	return reactor
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

func normalise(i []instruction) ([]instruction, []int) {
	min := []int{i[0].Coords[0][0], i[0].Coords[1][0], i[0].Coords[2][0]}
	max := []int{i[0].Coords[0][0], i[0].Coords[1][0], i[0].Coords[2][0]}

	for index, instr := range i {
		for c, coord := range instr.Coords {
			// x1 > x1 swap
			if coord[0] > coord[1] {
				i[index].Coords[c] = [2]int{coord[1], coord[0]}
			}

			// normalisation
			if i[index].Coords[c][0] < min[c] {
				min[c] = i[index].Coords[c][0]
			}
			if i[index].Coords[c][1] > max[c] {
				max[c] = i[index].Coords[c][1]
			}
		}
	}

	// move min to 0
	for i := range max {
		max[i] = max[i] - min[i]
	}

	// shift values to min = 0
	for index, instr := range i {
		for c, coord := range instr.Coords {
			i[index].Coords[c][0] = coord[0] - min[c]
			i[index].Coords[c][1] = coord[1] - min[c]
		}
	}

	return i, max
}

func rebootAndCount(i []instruction) int {
	fmt.Printf("Noramlise\n")
	instructions, size := normalise(i)
	// normalise
	// calc overlap
	// calc volume
	fmt.Printf("Bootup\n")
	reactor := bootUp(instructions, size)
	fmt.Printf("Count\n")
	numOfCubes := countCubes(reactor)
	fmt.Printf("Done\n")
	return numOfCubes
}
