package divide_chocolate

import "math"

//Runtime: 14 ms, faster than 100.00% of Go online submissions for Divide Chocolate.
//Memory Usage: 6.2 MB, less than 100.00% of Go online submissions for Divide Chocolate.
func maximizeSweetness(sweetness []int, k int) int {
	l, r := minSumTuple(sweetness)
	for l < r {
		m := (l + r) >> 1
		if feasible(sweetness, m, k) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func feasible(sweetness []int, m, k int) bool {
	count, total := 0, 0
	for _, sweet := range sweetness {
		total += sweet
		if total > m {
			count++
			total = 0
		}
	}
	return count <= k
}

func minSumTuple(arr []int) (int, int) {
	min, sum := math.MaxInt, 0
	for _, item := range arr {
		if min > item {
			min = item
		}
		sum += item
	}
	return min, sum
}
