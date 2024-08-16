package main

import "fmt"

func buildMap(replacements [][]string) map[string]string {
	newMap := map[string]string{}
	for i := range replacements {
		newMap[replacements[i][0]] = replacements[i][1]
	}
	return newMap
}

func calcPolyValue(letters string, replacements [][]string) int {
	replaceMap := buildMap(replacements)

	newString := string(letters[0])
	for k := 0; k < 10; k++ {

		fmt.Println(k, newString)
		for i := range letters[:len(letters)-1] {
			x := string(letters[i]) + string(letters[i+1])
			newString += replaceMap[x]
			newString += string(letters[i+1])
		}
	}

	/*
		for i := range replacements {
			from := replacements[i][0]

			to := string(from[0]) + replacements[i][1] + string(from[1])
			s = strings.Replace(s, from, to, 1)
			fmt.Println(from, to, s)
		}
		fmt.Println(s)
	*/
	return 0
}
