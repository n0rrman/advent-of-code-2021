package main

type dice int

func (d *dice) rollDice() int {
	*d = ((*d)%100 + 1)
	return int(*d)
}

func (d *dice) readDice() int {
	return int(*d)
}
