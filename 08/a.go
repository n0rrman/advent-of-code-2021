package main

func calc1478(out [][]digit) int {
	total := 0

	for _, digits := range out {
		for _, row := range digits {

			count := 0
			for _, bit := range row {
				if bit {
					count++
				}
			}

			if count == 2 || count == 4 || count == 3 || count == 7 {
				total++
			}
		}
	}
	return total
}
