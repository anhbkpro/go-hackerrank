package max_consecutive_ones

func FindMaxConsecutiveOnes(nums []int) int {
	l, ans := 0, 0
	freqWindow := make(map[int]int)
	for r := 0; r < len(nums); r++ {
		freqWindow[nums[r]]++
		for freqWindow[0] != 0 {
			freqWindow[nums[l]]--
			if freqWindow[nums[l]] == 0 {
				delete(freqWindow, nums[l])
			}
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}

func FindMaxConsecutiveOnesV2(nums []int) int {
	maxHere, max := 0, 0
	for _, num := range nums {
		if num == 0 {
			maxHere = 0
		} else {
			maxHere++
		}
		if maxHere > max {
			max = maxHere
		}
	}
	return max
}
