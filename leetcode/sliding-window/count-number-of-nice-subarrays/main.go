package count_number_of_nice_subarrays

func NumberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	ans := 0
	left := 0
	right := 0
	count := 0
	for right < n {
		if nums[right]%2 == 1 {
			k--
			count = 0
		}

		for k == 0 {
			k += nums[left] % 2
			left++
			count++
		}

		ans += count
		right++
	}

	return ans
}
