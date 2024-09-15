package main

import "fmt"

type scanner struct {
	beacons   [][]int
	distances map[[2]int]float64
}

func (s *scanner) init(beacons [][]int) {
	copy(s.beacons, beacons)
}

func (s *scanner) calcDistances() {
	// Calculate 3d distances
}

func syncAndCountScanners(scanners [][][]int) int {
	allScanners := make([]scanner, len(scanners))

	for i, s := range scanners {
		fmt.Println(i)
		newScanner := scanner{}
		newScanner.init(s)
		newScanner.calcDistances()
		allScanners[i] = newScanner
	}
	fmt.Println(allScanners)
	//fmt.Println(beacons[1], scanners[3][1])
	return 0
}
