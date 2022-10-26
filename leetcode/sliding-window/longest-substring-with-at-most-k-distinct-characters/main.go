package longest_substring_with_at_most_k_distinct_characters

import "math"

func LongestKSubstr(s string, k int) int {
	l, r, max := 0, 0, 0
	freqWindow := make(map[string]int)
	for r < len(s) {
		cRight := string(s[r])
		if _, ok := freqWindow[cRight]; ok {
			freqWindow[cRight]++
		} else {
			freqWindow[cRight] = 1
		}
		if len(freqWindow) <= k {
			max = int(math.Max(float64(max), float64(r-l+1)))
		}
		if len(freqWindow) <= k {
			r++
		} else {
			cLeft := string(s[l])
			freqWindow[cLeft]--
			if freqWindow[cLeft] == 0 {
				delete(freqWindow, cLeft)
			}
			r++
			l++
		}
	}
	return max
}
