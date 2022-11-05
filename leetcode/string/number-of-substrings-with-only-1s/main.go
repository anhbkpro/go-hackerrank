package number_of_substrings_with_only_1s

const MOD = 1000000007

// NumSub Similar questions: https://leetcode.com/problems/count-number-of-homogenous-substrings/
func NumSub(s string) int {
	ans := 0
	var cur rune
	runes := []rune(s)
	count := 0
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == "0" {
			count = 0
			continue
		}
		if runes[i] == cur {
			count++
		} else {
			count = 1
		}
		cur = runes[i]
		ans = (ans + count) % MOD
	}
	return ans
}
