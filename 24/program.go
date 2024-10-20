package main

type instructionType int
type program []instruction

const (
	Inp instructionType = iota + 1
	Add
	Mul
	Div
	Mod
	Eql
)

type instruction struct {
	instr  instructionType
	varA   *int
	varB   *int
	constA int
	constB int
}

type variables struct {
	w int
	x int
	y int
	z int
}
