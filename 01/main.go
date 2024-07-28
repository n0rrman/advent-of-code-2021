package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) []string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	data := strings.Split(string(body[:]), "\n")
	return data
}

func calcIncr(startVal string, vals []string) int {
	prev := startVal
	count := 0
	for _, d := range vals {

		a, _ := strconv.Atoi(d)
		b, _ := strconv.Atoi(prev)
		if a > b {
			count += 1
		}
		prev = d
	}

	return count
}

func main() {
	test_data := readData("test_data")
	data := readData("data")

	test := calcIncr(test_data[0], test_data[1:])
	results := calcIncr(data[0], data[1:])
	if test != 7 {
		os.Exit(1)
	}

	fmt.Println(results)
}
