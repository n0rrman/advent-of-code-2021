package main

import "fmt"

func printSheet(sheet [][]bool) {
	for i := range sheet {
		for j := range sheet[i] {
			if sheet[i][j] {
				fmt.Print("@ ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func foldAndPrint(points [][]int, folds [][]int) {
	fullSheet := buildSheet(points)

	for i := range folds {
		fullSheet = foldSheet(fullSheet, folds[i])
	}

	printSheet(fullSheet)
}
