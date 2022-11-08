package count_number_of_nice_subarrays

func NumberOfSubarrays(nums []int, k int) int {
	l, r, ans, count, n := 0, 0, 0, 0, len(nums)
	for r < n {
		if nums[r]%2 == 1 {
			k--
			count = 0
		}

		for k == 0 {
			k += nums[l] % 2
			l++
			count++
		}

		ans += count
		r++
	}

	return ans
}
