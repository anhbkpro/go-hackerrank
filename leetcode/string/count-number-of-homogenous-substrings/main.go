package count_number_of_homogenous_substrings

const MOD = 1000000007

func CountHomogenous(s string) int {
	runes := []rune(s)
	ans := 0
	var cur rune
	count := 0
	for i := 0; i < len(runes); i++ {
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
