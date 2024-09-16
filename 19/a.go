package main

import (
	"fmt"
	"math"
)

type scanner struct {
	beacons   [][]int
	distances map[[2]int]float64
}

func (s *scanner) init(beacons [][]int) {
	(*s).beacons = make([][]int, len(beacons))
	for i, b := range beacons {
		(*s).beacons[i] = make([]int, 3)
		copy((*s).beacons[i], b)
	}

	(*s).calcDistances()
}

func (s *scanner) calcDistances() {
	(*s).distances = make(map[[2]int]float64)
	for i := 0; i < len(s.beacons); i++ {
		for j := i + 1; j < len(s.beacons); j++ {
			(*s).distances[[2]int{i, j}] = calc3dDistance(s.beacons[i], s.beacons[j])
		}
	}
}

func calc3dDistance(a []int, b []int) float64 {
	tmp := math.Pow(float64(a[0])-float64(b[0]), 2) + math.Pow(float64(a[1])-float64(b[1]), 2) + math.Pow(float64(a[2])-float64(b[2]), 2)
	return math.Sqrt(tmp)
}

func syncAndCountScanners(scanners [][][]int) int {
	allScanners := make([]scanner, len(scanners))

	for i, s := range scanners {
		newScanner := scanner{}
		newScanner.init(s)
		allScanners[i] = newScanner
	}
	fmt.Println(allScanners[0].beacons)
	fmt.Println(allScanners[0].distances)
	return 0
}
