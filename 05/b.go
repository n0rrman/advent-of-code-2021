package main

func calcVentOverlapB(coords []coordPair) int {
	var vm ventMap
	for _, c := range coords {
		vm = vm.extendMap(c)

		if c.fromX == c.toX {
			vm = vm.markVertical(c)
		} else if c.fromY == c.toY {
			vm = vm.markHorizontal(c)
		} else {
			if c.fromY < c.toY && c.fromX < c.toX {
				vm = vm.markDiagDown(c)
			}
			if c.fromY > c.toY && c.fromX > c.toX {
				vm = vm.markDiagDown(c)
			}
			if c.fromY < c.toY && c.fromX > c.toX {
				vm = vm.markDiagUp(c)
			}
			if c.fromY > c.toY && c.fromX < c.toX {
				vm = vm.markDiagUp(c)
			}
		}
	}
	return vm.countOverlap(2)
}
