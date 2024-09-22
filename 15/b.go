package main

func extendMap(data [][]int) [][]int {
	extendedMap := make([][]int, len(data)*5)

	for y := range extendedMap {
		extendedMap[y] = make([]int, len(data[0])*5)
		for x := range extendedMap[y] {
			extension := y/len(data) + x/len(data[0])
			extendedMap[y][x] = (((extension + data[y%(len(data))][x%(len(data[0]))]) - 1) % 9) + 1
		}
	}
	return extendedMap
}

func calcLowestRiskExtended(data [][]int) int {
	extendedMap := extendMap(data)
	return calcLowestRisk(extendedMap)
}
