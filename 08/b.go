package main

import (
	"strconv"
)

func all(val []bool) bool {
	for _, val := range val {
		if !val {
			return false
		}
	}
	return true
}

func (a digit) OR(b digit) digit {
	var c digit
	for i := range c {
		c[i] = a[i] || b[i]
	}

	return c
}

func (a digit) XOR(b digit) digit {
	var c digit
	for i := range c {
		c[i] = (a[i] || b[i]) && !(a[i] && b[i])
	}

	return c
}

func (a digit) has(b digit) bool {
	for i := range a {
		if b[i] && !a[i] {
			return false
		}
	}
	return true
}

func (d digit) len() int {
	count := 0
	for _, x := range d {
		if x {
			count++
		}
	}
	return count
}

func decode(d []digit) map[digit]int {

	digToVal := make(map[digit]int)
	valToDig := make(map[int]digit)

	for _, val := range d {
		if val.len() == 2 {
			digToVal[val] = 1
			valToDig[1] = val
		} else if val.len() == 4 {
			digToVal[val] = 4
			valToDig[4] = val
		} else if val.len() == 3 {
			digToVal[val] = 7
			valToDig[3] = val
		} else if val.len() == 7 {
			digToVal[val] = 8
			valToDig[8] = val
		}
	}

	for _, val := range d {
		if val.len() == 5 {
			if val.has(valToDig[4].XOR(valToDig[1])) {
				digToVal[val] = 5
				valToDig[5] = val
			} else if val.has(valToDig[1]) {
				digToVal[val] = 3
				valToDig[3] = val
			} else {
				digToVal[val] = 2
				valToDig[2] = val
			}
		}
		if val.len() == 6 {
			if val.has(valToDig[4]) {
				digToVal[val] = 9
				valToDig[9] = val
			} else if val.has(valToDig[8].XOR(valToDig[1])) {
				digToVal[val] = 6
				valToDig[6] = val
			} else {
				digToVal[val] = 0
				valToDig[0] = val
			}
		}
	}

	return digToVal

}

func calcSum(in [][]digit, out [][]digit) int {
	sum := 0

	for i := range in {
		numbers := decode(append(in[i], out[i]...))
		section := ""

		for _, num := range out[i] {
			numString := strconv.Itoa(numbers[num])
			section += numString
		}
		sectionVal, _ := strconv.Atoi(section)
		sum += sectionVal
	}

	return sum
}
