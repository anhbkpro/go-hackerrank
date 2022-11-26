package cutting_ribbons

import "math"

//Runtime: 420 ms, faster than 65.00% of Go online submissions for Cutting Ribbons.
//Memory Usage: 8.8 MB, less than 85.00% of Go online submissions for Cutting Ribbons.
func maxLength(ribbons []int, k int) int {
	l := 1
	r, total := maxSumTuple(ribbons)
	r++            // everytime you want to find the maximum value --> upper bound need +1
	if total < k { // cannot obtain k ribbons of the same positive integer length.
		return 0
	}

	for l < r {
		m := (l + r) >> 1
		if !isCutPossible(ribbons, k, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}

func isCutPossible(ribbons []int, k, length int) bool {
	total := 0
	for _, ribbon := range ribbons {
		total += ribbon / length
	}
	return k <= total
}

func maxSumTuple(arr []int) (int, int) {
	res, sum := math.MinInt, 0
	for _, item := range arr {
		sum += item
		if res < item {
			res = item
		}
	}
	return res, sum
}
