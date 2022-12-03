package longest_palindromic_substring

func LongestPalindrome(s string) string {
	ans := ""
	ansLen := 0
	for i := 0; i < len(s); i++ {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > ansLen {
				ans = s[l : r+1]
				ansLen = r - l + 1
			}
			l -= 1
			r += 1
		}

		// even length
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > ansLen {
				ans = s[l : r+1]
				ansLen = r - l + 1
			}
			l -= 1
			r += 1
		}
	}
	return ans
}

//Runtime: 51 ms, faster than 72.73% of Go online submissions for Longest Palindromic Subsequence.
//Memory Usage: 20.2 MB, less than 34.54% of Go online submissions for Longest Palindromic Subsequence.
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// Base case: Single letter substrings
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
