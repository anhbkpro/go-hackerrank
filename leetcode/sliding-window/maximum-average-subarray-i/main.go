package maximum_average_subarray_i

import "math"

func FindMaxAverage(nums []int, k int) float64 {
	windowSum, maxSum, l := 0, math.MinInt, 0
	for r := 0; r < len(nums); r++ {
		windowSum += nums[r]
		if r-l+1 == k {
			maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
			windowSum -= nums[l]
			l++
		}
	}
	return float64(maxSum) / float64(k)
}

func FindMaxAverageV2(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		max = int(math.Max(float64(max), float64(sum)))
	}
	return float64(max) / float64(k)
}
