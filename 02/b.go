package main

func calcPosAndAim(instructions []instruction) int {
	x := 0
	y := 0
	aim := 0
	for _, i := range instructions {
		switch i.method {
		case "up":
			aim -= i.count
		case "down":
			aim += i.count
		case "forward":
			x += i.count
			y += aim * i.count
		default:
			x -= i.count
			x -= aim * i.count
		}
	}
	return x * y
}
