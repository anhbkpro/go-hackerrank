package maximum_erasure_value

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
		intSet[nums[r]] = true
		if windowSum > ans { // maybe of test suit, same performance as math.Max
			ans = windowSum
		}
	}
	return ans
}
