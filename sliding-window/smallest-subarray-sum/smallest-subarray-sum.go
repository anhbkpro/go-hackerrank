package smallest_subarray_sum

import "math"

func MinSizeSubArraySum(s int32, arr []int32) int32 {
	minLength := int32(math.MaxInt32)
	var windowStart int32
	var windowSum int32
	for windowEnd := int32(0); windowEnd < int32(len(arr)); windowEnd++ {
		windowSum += arr[windowEnd]
		for windowSum >= s {
			minLength = int32(math.Min(float64(minLength), float64(windowEnd-windowStart+1)))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return minLength
}
