package main

import "fmt"

var rollMultiplier = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func recPlayAndCount(posA int, posB int, scoreA int, scoreB int) (int, int) {
	rolls := []int{3, 4, 5, 6, 7, 8, 9}
	if scoreB >= 21 {
		return 0, 1
	}
	if scoreA >= 21 {
		return 1, 0
	}

	winsA, winsB := 0, 0
	for _, rollA := range rolls {
		for _, rollB := range rolls {

			posA = ((posA + rollA - 1) % 10) + 1
			scoreA += posA

			if scoreA < 21 {

				a, b := recPlayAndCount(posA, posB, scoreA, scoreB)
				winsA += a * rollMultiplier[rollA]
				winsB += b * rollMultiplier[rollB]

			}

			posB = ((posB + rollB - 1) % 10) + 1
			scoreB += posB

			if scoreB < 21 {

				a, b := recPlayAndCount(posA, posB, scoreA, scoreB)
				winsA += a * rollMultiplier[rollA]
				winsB += b * rollMultiplier[rollB]
			}
		}
	}
	return winsA, winsB
}

func playProperlyAndCount(posA int, posB int) int {
	aWins, bWins := recPlayAndCount(posA, posB, 0, 0)

	fmt.Println(aWins, bWins)

	if aWins > bWins {
		return aWins
	} else {
		return bWins
	}

}
