package max_subarrays

import "math"

func MaxSumSubArrayOfSizeK(k int32, arr []int32) int32 {
	var result, windowStart, windowSum int32
	for windowEnd := int32(0); windowEnd < int32(len(arr)); windowEnd++ {
		windowSum += arr[windowEnd]
		if windowEnd >= k-1 {
			result = int32(math.Max(float64(result), float64(windowSum)))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return result
}
