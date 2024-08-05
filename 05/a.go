package main

func calcVentOverlapA(coords []coordPair) int {
	var vm ventMap
	for _, c := range coords {
		vm = vm.extendMap(c)
		if c.fromX == c.toX {
			vm = vm.markVertical(c)
		} else if c.fromY == c.toY {
			vm = vm.markHorizontal(c)
		}
	}

	return vm.countOverlap(2)
}
