package k_radius_subarray_averages

func GetAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	if 1+2*k > len(nums) {
		for i := 0; i < len(nums); i++ {
			ans[i] = -1
		}
		return ans
	}

	if k == 0 {
		return nums
	}

	l, windowSum := 0, 0
	for r := 0; r < len(nums); r++ {
		windowSum += nums[r]
		if r < k || len(nums)-1-r < k {
			ans[r] = -1
		}
		if r-l+1 == 1+2*k {
			ans[r-k] = windowSum / (2*k + 1)
			windowSum -= nums[l]
			l++
		}
	}

	return ans
}
