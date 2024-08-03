package main

import (
	"math"
)

type workList []int

func (wl workList) add(v int) workList {
	return append(wl, v)
}

func (wl workList) next() (int, workList) {
	if len(wl) > 0 {
		return wl[0], wl[1:]
	} else {
		return -1, workList{}
	}
}

func bitToInt(b bits) int {
	val := 0
	for i, bit := range b {
		val += bit * int(math.Pow(2, float64((len(b)-1-i))))
	}
	return val
}

func calcOxygen(b []bits) int {

	var wl workList
	for i := range b {
		wl = wl.add(i)
	}

	for bit := range b[0] {
		var (
			zero workList
			one  workList
		)
		for i, w := wl.next(); i >= 0; i, w = w.next() {
			if b[i][bit] == 0 {
				zero = zero.add(i)
			} else {
				one = one.add(i)
			}
		}

		if len(one) >= len(zero) {
			wl = one
		} else {
			wl = zero
		}
	}
	return bitToInt(b[wl[0]])
}

func calcCO2(b []bits) int {
	var wl workList
	for i := range b {
		wl = wl.add(i)
	}

	for bit := range b[0] {
		if len(wl) == 1 {
			break
		}

		var (
			zero workList
			one  workList
		)
		for i, w := wl.next(); i >= 0; i, w = w.next() {
			if b[i][bit] == 0 {
				zero = zero.add(i)
			} else {
				one = one.add(i)
			}
		}

		if len(one) >= len(zero) {
			wl = zero
		} else {
			wl = one
		}
	}
	return bitToInt(b[wl[0]])
}

func calcLSR(b []bits) int {
	o := calcOxygen(b)
	co2 := calcCO2(b)

	return o * co2
}
