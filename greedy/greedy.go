package greedy

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, boxType := range boxTypes {
		if boxType[0] > truckSize {
			continue
		}
		ans += boxType[0] * boxType[1]
		truckSize -= boxType[0]

	}
	return
}
