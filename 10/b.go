package main

import (
	"sort"
)

func calcAutocompScore(brackets [][]string) int {

	var scores []int

	for i := range brackets {
		var stack []string
		sum := 0
		skip := false
		for _, val := range brackets[i] {
			if val == "{" || val == "(" || val == "<" || val == "[" {
				stack = append(stack, val)
			} else {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if flipBracket(pop) != val {
					skip = true
					break
				}
			}
		}
		if !skip {
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == "(" {
					sum = 5*sum + 1
				} else if stack[i] == "[" {
					sum = 5*sum + 2
				} else if stack[i] == "{" {
					sum = 5*sum + 3
				} else if stack[i] == "<" {
					sum = 5*sum + 4
				}
			}
			scores = append(scores, sum)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}
