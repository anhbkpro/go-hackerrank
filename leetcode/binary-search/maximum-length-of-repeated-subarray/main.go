package maximum_length_of_repeated_subarray

//Runtime: 79 ms, faster than 58.00% of Go online submissions for Maximum Length of Repeated Subarray.
//Memory Usage: 20.7 MB, less than 10.00% of Go online submissions for Maximum Length of Repeated Subarray.
func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	SIZE1 := len(nums1) + 1
	SIZE2 := len(nums2) + 1
	dp := make([][]int, SIZE1+1)
	for i := 0; i < SIZE1+1; i++ {
		dp[i] = make([]int, SIZE2+1)
	}
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans
}
