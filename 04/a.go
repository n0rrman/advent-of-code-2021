package main

func calcFirstWin(randomNums []int, boards [][][]int) int {
	boardList := make([]board, len(boards))

	for i := range boardList {
		boardList[i].new(boards[i])
	}

	for _, num := range randomNums {
		for _, b := range boardList {
			b.play(num)
			if b.hasWon() {
				return b.getScore() * b.lastPlayed
			}
		}
	}
	return 0
}
