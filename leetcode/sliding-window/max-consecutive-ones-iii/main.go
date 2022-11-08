package max_consecutive_ones_iii

func LongestOnes(nums []int, k int) int {
	l, ans := 0, 0
	freqWindow := make(map[int]int)
	for r := 0; r < len(nums); r++ {
		freqWindow[nums[r]]++
		for freqWindow[0] > k {
			freqWindow[nums[l]]--
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}

// LongestOnesV2 Runtime: 57 ms, faster than 94.19% of Go online submissions for Max Consecutive Ones III.
func LongestOnesV2(nums []int, k int) int {
	l, r, numZeroes, ans := 0, 0, 0, 0
	for r = 0; r < len(nums); r++ {
		if nums[r] == 0 {
			numZeroes++
		}
		if numZeroes > k {
			if nums[l] == 0 {
				numZeroes--
			}
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}
