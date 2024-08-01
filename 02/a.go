package main

func calcPos(instructions []instruction) int {
	x := 0
	y := 0
	for _, i := range instructions {
		switch i.method {
		case "up":
			y -= i.count
		case "down":
			y += i.count
		case "forward":
			x += i.count
		default:
			x -= i.count
		}
	}
	return x * y
}
