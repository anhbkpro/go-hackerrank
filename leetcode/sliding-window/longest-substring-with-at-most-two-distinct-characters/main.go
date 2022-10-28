package longest_substring_with_at_most_two_distinct_characters

import "math"

func LengthOfLongestSubstringTwoDistinct(s string) int {
	l, r, max := 0, 0, 0
	freqWindow := make(map[string]int)
	for r < len(s) {
		cRight := string(s[r])
		if _, ok := freqWindow[cRight]; ok {
			freqWindow[cRight]++
		} else {
			freqWindow[cRight] = 1
		}

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
