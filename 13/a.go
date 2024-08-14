package main

func buildSheet(points [][]int) [][]bool {
	maxX := 0
	maxY := 0
	for i := range points {
		x := points[i][0]
		y := points[i][1]
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}

	sheet := make([][]bool, maxY+1)
	for i := range sheet {
		sheet[i] = make([]bool, maxX+1)
	}

	for i := range points {
		x := points[i][0]
		y := points[i][1]
		sheet[y][x] = true
	}
	return sheet
}

func foldUp(sheet [][]bool, foldY int) [][]bool {
	newSheet := make([][]bool, foldY)
	for i := range newSheet {
		newSheet[foldY-1-i] = make([]bool, len(sheet[i]))
		for j := range sheet[i] {
			newSheet[foldY-1-i][j] = false
			if foldY-1-i >= 0 {
				newSheet[foldY-1-i][j] = newSheet[foldY-1-i][j] || sheet[foldY-1-i][j]
			}
			if foldY+1+i < len(sheet) {
				newSheet[foldY-1-i][j] = newSheet[foldY-1-i][j] || sheet[foldY+1+i][j]
			}
		}
	}
	return newSheet
}

func foldLeft(sheet [][]bool, foldX int) [][]bool {
	newSheet := make([][]bool, len(sheet))
	for i := range newSheet {
		newSheet[i] = make([]bool, foldX)
		for j := range newSheet[i] {
			newSheet[i][foldX-j-1] = false
			if foldX-1-j >= 0 {
				newSheet[i][foldX-j-1] = newSheet[i][foldX-j-1] || sheet[i][foldX-j-1]
			}
			if foldX+1+j < len(sheet[i]) {
				newSheet[i][foldX-j-1] = newSheet[i][foldX-j-1] || sheet[i][foldX+1+j]
			}
		}
	}
	return newSheet
}

func foldSheet(sheet [][]bool, fold []int) [][]bool {
	if fold[0] == 0 {
		return foldLeft(sheet, fold[1])
	} else {
		return foldUp(sheet, fold[1])
	}

}

func countMarks(sheet [][]bool) int {
	count := 0

	for y := range sheet {
		for x := range sheet[y] {
			if sheet[y][x] {
				count++
			}
		}
	}

	return count
}

func foldAndCount(points [][]int, folds [][]int, numFolds int) int {
	fullSheet := buildSheet(points)

	for i := 0; i < numFolds; i++ {
		fullSheet = foldSheet(fullSheet, folds[i])
	}
	return countMarks(fullSheet)
}
