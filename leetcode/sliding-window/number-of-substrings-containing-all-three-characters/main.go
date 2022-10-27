package number_of_substrings_containing_all_three_characters

func NumberOfSubstrings(s string) int {
	l, r, ans := 0, 0, 0
	freqWindow := make(map[string]int)
	for r < len(s) {
		cRight := string(s[r])
		if _, ok := freqWindow[cRight]; ok {
			freqWindow[cRight]++
		} else {
			freqWindow[cRight] = 1
		}
		for freqWindow["a"] > 0 && freqWindow["b"] > 0 && freqWindow["c"] > 0 {
			cLeft := string(s[l])
			freqWindow[cLeft]--
			l++
		}
		ans += l
		r++
	}
	return ans
}
