package main

import "fmt"

func buildMap(replacements [][]string) map[string]string {
	newMap := map[string]string{}
	for i := range replacements {
		newMap[replacements[i][0]] = replacements[i][1]
	}
	return newMap
}

func calcScore(letters string) int {
	counts := make([]int, int('Z'))

	for _, letter := range letters {
		counts[int(letter)] += 1
	}

	max := 0
	min := len(letters)

	for _, val := range counts {
		if val > max {
			max = val
		}
		if val < min && val != 0 {
			min = val
		}
	}
	fmt.Println(max, min)
	score := max - min
	return score
}

func calcPolyValue(letters string, replacements [][]string, iter int) int {
	replaceMap := buildMap(replacements)

	str := letters

	for k := 1; k <= iter; k++ {
		newString := string(letters[0])
		for i := range str[:len(str)-1] {
			x := string(str[i]) + string(str[i+1])
			newString += replaceMap[x]
			newString += string(str[i+1])
		}
		str = newString
	}

	return calcScore(str)
}
