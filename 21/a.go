package main

func playAndCount(posA int, posB int) int {
	var d dice = 0

	rolled := 0
	scoreA := 0
	scoreB := 0

	for {
		posA = ((posA + d.rollDice() + d.rollDice() + d.rollDice() - 1) % 10) + 1
		scoreA += posA
		rolled += 3
		if scoreA >= 1000 {
			break
		}

		posB = ((posB + d.rollDice() + d.rollDice() + d.rollDice() - 1) % 10) + 1
		scoreB += posB
		rolled += 3
		if scoreB >= 1000 {
			break
		}

	}
	if scoreA > scoreB {
		return scoreB * rolled
	} else {
		return scoreA * rolled
	}
}
