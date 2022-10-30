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
	left, right, numZeroes, ans := 0, 0, 0, 0
	for right = 0; right < len(nums); right++ {
		if nums[right] == 0 {
			numZeroes++
		}
		if numZeroes > k {
			if nums[left] == 0 {
				numZeroes--
			}
			left++
		}
		if numZeroes <= k {
			if right-left+1 > ans {
				ans = right - left + 1
			}
		}
	}
	return ans
}
