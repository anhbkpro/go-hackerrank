package heaters

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Slice(heaters, func(i, j int) bool {
		return heaters[i] < heaters[j]
	})
	res := math.MinInt

	for _, house := range houses {
		index := sort.Search(len(heaters), func(i int) bool {
			return heaters[i] >= house
		})
		
		if index < 0 {
			index = -(index + 1)
		}

		dist1, dist2 := 0, 0
		if index-1 >= 0 {
			dist1 = house - heaters[index-1]
		} else {
			dist1 = math.MaxInt
		}
		if index < len(heaters) {
			dist2 = heaters[index] - house
		} else {
			dist2 = math.MaxInt
		}
		res = max(res, min(dist1, dist2))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
