package minimum_size_subarray_sum

import "math"

func MinSubArrayLen(target int, nums []int) int {
	l, ans, windowSum := 0, math.MaxInt, 0
	for r := 0; r < len(nums); r++ {
		windowSum += nums[r]
		for windowSum >= target {
			ans = int(math.Min(float64(ans), float64(r-l+1)))
			windowSum -= nums[l]
			l++
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}
