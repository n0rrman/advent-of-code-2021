package main

func flipBracket(b string) string {
	brackets := make(map[string]string)
	brackets["("] = ")"
	brackets["["] = "]"
	brackets["{"] = "}"
	brackets["<"] = ">"
	brackets[")"] = "("
	brackets["]"] = "["
	brackets["}"] = "{"
	brackets[">"] = "<"

	return brackets[b]
}

func getValue(b string) int {
	brackets := make(map[string]int)
	brackets[")"] = 3
	brackets["]"] = 57
	brackets["}"] = 1197
	brackets[">"] = 25137

	return brackets[b]
}

func calcSysErr(brackets [][]string) int {
	sum := 0

	for i := range brackets {
		var stack []string
		for _, val := range brackets[i] {
			if val == "{" || val == "(" || val == "<" || val == "[" {
				stack = append(stack, val)
			} else {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if flipBracket(pop) != val {
					sum += getValue(val)
					break
				}
			}
		}
	}

	return sum
}
