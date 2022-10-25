package minimum_size_subarray_sum

import "math"

func MinSubArrayLen(target int, nums []int) int {
	l, r, min, windowSum := 0, 0, math.MaxInt, 0
	for r < len(nums) {
		windowSum += nums[r]
		for windowSum >= target {
			windowSum -= nums[l]
			min = int(math.Min(float64(min), float64(r-l+1)))
			l++
		}
		r++
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}
