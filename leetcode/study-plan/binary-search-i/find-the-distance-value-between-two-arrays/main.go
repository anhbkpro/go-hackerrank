package find_the_distance_value_between_two_arrays

import (
	"math"
	"sort"
)

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	ans := 0
	for _, num := range arr1 {
		if isValid(arr2, num, d) {
			ans++
		}
	}
	return ans
}

func isValid(arr []int, target, diff int) bool {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if int(math.Abs(float64(arr[mid]-target))) <= diff {
			return false
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return true
}
