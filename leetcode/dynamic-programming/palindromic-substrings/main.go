package palindromic_substrings

//Runtime: 6 ms, faster than 34.34% of Go online submissions for Palindromic Substrings.
//Memory Usage: 4.6 MB, less than 17.47% of Go online submissions for Palindromic Substrings.
func countSubstrings(s string) int {
	n, ans := len(s), 0
	dp := make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n)
		ans++
	}
	// Base case: Single letter substrings
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Base case: double letter substrings
	for i := 0; i < n-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			ans++
		}
	}

	// All other cases: substrings of length 3 to n
	i := 0
	for len := 3; len <= n; len++ {
		for i, j := 0, i+len-1; j < n; i, j = i+1, j+1 {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] {
				ans++
			}
		}
	}

	return ans
}
