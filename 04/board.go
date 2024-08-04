package main

type board struct {
	marks      [][]bool
	numbers    [][]int
	lastPlayed int
	score      int
}

func (b *board) play(num int) bool {
	if b.score != 0 {
		return false
	}
	b.lastPlayed = num

	for y, cols := range b.numbers {
		for x, val := range cols {
			if val == num {
				b.marks[y][x] = true
				return true
			}
		}
	}
	return true
}

func (b *board) new(values [][]int) {
	b.score = 0
	b.marks = make([][]bool, len(values))
	b.numbers = make([][]int, len(values))

	for y, cols := range values {
		markRow := make([]bool, len(values[0]))
		numberRow := make([]int, len(values[0]))
		for x, val := range cols {
			markRow[x] = false
			numberRow[x] = val
		}
		b.marks[y] = markRow
		b.numbers[y] = numberRow
	}
}

func (b *board) getScore() int {

	if b.score != 0 {
		return b.score
	}

	for y, cols := range b.marks {
		for x, marked := range cols {
			if !marked {
				b.score += b.numbers[y][x]
			}
		}
	}
	return b.score
}

func (b *board) hasWon() bool {
	if b.score != 0 {
		return false
	}
	winning := false
	for _, rows := range b.marks {
		for _, mark := range rows {
			winning = mark
			if !winning {
				break
			}
		}
		if winning {
			return true
		}
	}

	for rows := range b.marks[0] {
		for cols := range b.marks {
			winning = b.marks[cols][rows]
			if !winning {
				break
			}
		}
		if winning {
			return true
		}

	}
	return false
}
