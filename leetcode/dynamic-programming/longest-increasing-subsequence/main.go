package longest_increasing_subsequence

//	Dynamic Programming:
//		Runtime: 71 ms, faster than 48.37% of Go online submissions for Longest Increasing Subsequence.
//		Memory Usage: 3.8 MB, less than 39.85% of Go online submissions for Longest Increasing Subsequence.
//	Binary Search:
//		Runtime: 5 ms, faster than 97.49% of Go online submissions for Longest Increasing Subsequence.
//		Memory Usage: 3.6 MB, less than 64.91% of Go online submissions for Longest Increasing Subsequence.
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i <= len(nums)-1; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	longest := 0
	for _, c := range dp {
		longest = max(longest, c)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
