package number_of_substrings_with_only_1s

const MOD = 1000000007

// NumSub Similar questions: https://leetcode.com/problems/count-number-of-homogenous-substrings/
func NumSub(s string) int {
	count := 0
	ans := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == "1" {
			count++
		} else {
			count = 0
		}
		ans = (ans + count) % MOD
	}
	return ans
}
