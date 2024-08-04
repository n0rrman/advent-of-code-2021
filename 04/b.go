package main

func calcLastWin(randomNums []int, boards [][][]int) int {
	boardList := make([]board, len(boards))

	for i := range boardList {
		boardList[i].new(boards[i])
	}

	lastWinner := 0
	for _, num := range randomNums {
		for i := range boardList {
			boardList[i].play(num)
			if boardList[i].hasWon() {
				lastWinner = i
				boardList[i].getScore()
			}
		}
	}
	b := boardList[lastWinner]
	return b.getScore() * b.lastPlayed
}
