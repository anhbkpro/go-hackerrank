package minimum_size_subarray_sum

import "math"

func MinSubArrayLen(target int, nums []int) int {
	l, min, windowSum := 0, math.MaxInt, 0
	for r := 0; r < len(nums); r++ {
		windowSum += nums[r]
		for windowSum >= target {
			min = int(math.Min(float64(min), float64(r-l+1)))
			windowSum -= nums[l]
			l++
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}
