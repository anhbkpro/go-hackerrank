package subarrays_with_k_different_integers

func SubarraysWithKDistinct(nums []int, k int) int {
	return countsOfSubarraysWithKDistinct(nums, k) - countsOfSubarraysWithKDistinct(nums, k-1)
}

func countsOfSubarraysWithKDistinct(nums []int, k int) int {
	l, r, ans := 0, 0, 0
	freqWindow := make(map[int]int)
	for r < len(nums) {
		if _, ok := freqWindow[nums[r]]; ok {
			freqWindow[nums[r]]++
		} else {
			freqWindow[nums[r]] = 1
		}
		for len(freqWindow) > k {
			freqWindow[nums[l]]--
			if freqWindow[nums[l]] == 0 {
				delete(freqWindow, nums[l])
			}
			l++
		}
		ans += r - l + 1
		r++
	}
	return ans
}
