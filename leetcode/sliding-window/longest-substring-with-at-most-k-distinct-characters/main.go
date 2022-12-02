package longest_substring_with_at_most_k_distinct_characters

import "math"

// LongestKSubstr Runtime: 8 ms, faster than 82.86% of Go online submissions for Longest Substring with At Most K Distinct Characters.
//Memory Usage: 3.2 MB, less than 8.57% of Go online submissions for Longest Substring with At Most K Distinct Characters.
func LongestKSubstr(s string, k int) int {
	l, r, max := 0, 0, 0
	freqWindow := make(map[string]int)
	for r < len(s) {
		cRight := string(s[r])
		freqWindow[cRight]++
		if len(freqWindow) <= k {
			max = int(math.Max(float64(max), float64(r-l+1)))
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

func LongestKSubstrV2(s string, k int) int {
	l, max := 0, 0
	freqWindow := make(map[string]int)
	for r := 0; r < len(s); r++ {
		cRight := string(s[r])
		freqWindow[cRight]++

		for len(freqWindow) > k {
			cLeft := string(s[l])
			freqWindow[cLeft]--
			if freqWindow[cLeft] == 0 {
				delete(freqWindow, cLeft)
			}
			l++
		}
		max = int(math.Max(float64(max), float64(r-l+1)))
	}
	return max
}
