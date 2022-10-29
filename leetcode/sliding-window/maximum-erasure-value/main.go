package maximum_erasure_value

import "math"

func MaximumUniqueSubarray(nums []int) int {
	l, ans, windowSum := 0, 0, 0
	intSet := make(map[int]bool)
	for r := 0; r < len(nums); r++ {
		windowSum += nums[r]
		for intSet[nums[r]] {
			windowSum -= nums[l]
			delete(intSet, nums[l])
			l++
		}
		ans = int(math.Max(float64(ans), float64(windowSum)))
		intSet[nums[r]] = true
	}
	return ans
}
