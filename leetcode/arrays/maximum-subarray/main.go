package maximum_subarray

import "math"

func MaxSubArray(nums []int) int {
	largestSum := nums[0]
	windowSum := 0
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		if windowSum < 0 {
			windowSum = 0
		}
		windowSum += nums[windowEnd]
		largestSum = int(math.Max(float64(largestSum), float64(windowSum)))
	}

	return largestSum
}
