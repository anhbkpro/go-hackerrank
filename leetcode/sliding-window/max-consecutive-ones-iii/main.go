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
