package longest_substring_with_at_most_two_distinct_characters

import "math"

// LengthOfLongestSubstringTwoDistinct Runtime: 141 ms, faster than 22.50% of Go online submissions for Longest Substring with At Most Two Distinct Characters.
//Memory Usage: 6.6 MB, less than 17.50% of Go online submissions for Longest Substring with At Most Two Distinct Characters.
func LengthOfLongestSubstringTwoDistinct(s string) int {
	l, r, max := 0, 0, 0
	freqWindow := make(map[string]int)
	for r < len(s) {
		cRight := string(s[r])
		freqWindow[cRight]++
		if len(freqWindow) <= 2 {
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

// LengthOfLongestSubstringTwoDistinctV2 Runtime: 136 ms, faster than 25.00% of Go online submissions for Longest Substring with At Most Two Distinct Characters.
//Memory Usage: 6.7 MB, less than 15.00% of Go online submissions for Longest Substring with At Most Two Distinct Characters.
func LengthOfLongestSubstringTwoDistinctV2(s string) int {
	l, max := 0, 0
	freqWindow := make(map[string]int)
	for r := 0; r < len(s); r++ {
		cRight := string(s[r])
		freqWindow[cRight]++

		for len(freqWindow) > 2 {
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
