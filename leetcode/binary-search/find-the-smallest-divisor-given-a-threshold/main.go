package find_the_smallest_divisor_given_a_threshold

import "math"

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, max(nums)
	for l < r {
		m := (l + r) >> 1
		if feasible(nums, m, threshold) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func feasible(nums []int, m, threshold int) bool {
	totalDivisionResult := 0
	for _, num := range nums {
		totalDivisionResult += num / m
		if num%m != 0 {
			totalDivisionResult++
		}
	}
	return totalDivisionResult <= threshold
}

func max(arr []int) int {
	res := math.MinInt
	for _, i := range arr {
		if res < i {
			res = i
		}
	}
	return res
}
