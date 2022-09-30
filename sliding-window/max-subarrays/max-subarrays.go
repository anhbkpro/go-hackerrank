package max_subarrays

func MaxSumSubArrayOfSizeK(k int32, arr []int32) int32 {
	var result int32
	var windowStart int32
	var windowSum int32
	for windowEnd := int32(0); windowEnd < int32(len(arr)); windowEnd++ {
		windowSum += arr[windowEnd]
		if windowEnd >= k-1 {
			if windowSum > result {
				result = windowSum
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return result
}
