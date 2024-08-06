package main

type ventMap [][]int

func (vm ventMap) extendMap(c coordPair) ventMap {
	// Get current size
	yLen := len(vm)
	xLen := 0
	resized := false

	if yLen > 0 {
		xLen = len(vm[0])
	}

	// Check max X
	if c.fromX >= xLen {
		xLen = c.fromX + 1
		resized = true
	}
	if c.toX >= xLen {
		xLen = c.toX + 1
		resized = true
	}

	// Check max Y
	if c.fromY >= yLen {
		yLen = c.fromY + 1
		resized = true
	}
	if c.toY >= yLen {
		yLen = c.toY + 1
		resized = true
	}

	if !resized {
		return vm
	}

	newMap := make(ventMap, yLen)
	for i := range newMap {
		newMap[i] = make([]int, xLen)
		for j := range newMap[i] {
			if len(vm) > i && len(vm[0]) > j {
				newMap[i][j] = vm[i][j]
			} else {
				newMap[i][j] = 0
			}
		}
	}
	return newMap
}

func (vm ventMap) markHorizontal(c coordPair) ventMap {
	if c.fromX > c.toX {
		c.fromX, c.toX = c.toX, c.fromX
	}

	for i := c.fromX; i <= c.toX; i++ {
		vm[c.fromY][i] += 1
	}
	return vm
}

func (vm ventMap) markVertical(c coordPair) ventMap {
	if c.fromY > c.toY {
		c.fromY, c.toY = c.toY, c.fromY
	}

	for i := c.fromY; i <= c.toY; i++ {
		vm[i][c.fromX] += 1
	}
	return vm
}

func (vm ventMap) markDiagUp(c coordPair) ventMap {
	if c.fromX > c.toX {
		c.fromX, c.toX = c.toX, c.fromX
		c.fromY, c.toY = c.toY, c.fromY
	}
	step := 0

	for i := c.fromX; i <= c.toX; i++ {
		vm[c.fromY+step][i] += 1
		step--
	}
	return vm
}

func (vm ventMap) markDiagDown(c coordPair) ventMap {
	if c.fromX > c.toX {
		c.fromX, c.toX = c.toX, c.fromX
		c.fromY, c.toY = c.toY, c.fromY
	}

	step := 0
	for i := c.fromX; i <= c.toX; i++ {
		vm[c.fromY+step][i] += 1
		step++
	}
	return vm
}

func (vm ventMap) countOverlap(threshold int) int {
	count := 0
	for _, rows := range vm {
		for _, val := range rows {
			if val >= threshold {
				count += 1
			}
		}
	}
	return count
}
