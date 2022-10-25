package longest_substring_without_repeating_characters

import "math"

func LengthOfLongestSubstring(s string) int {
	windowCounts := make(map[string]int)
	l, r, max := 0, 0, 0
	for r < len(s) {
		c := string(s[r])
		if _, ok := windowCounts[c]; ok {
			windowCounts[c]++
		} else {
			windowCounts[c] = 1
		}
		if len(windowCounts) == r-l+1 {
			max = int(math.Max(float64(max), float64(r-l+1)))
			r++
		} else if len(windowCounts) < r-l+1 {
			for len(windowCounts) < r-l+1 {
				c = string(s[l])
				windowCounts[c]--
				if windowCounts[c] == 0 {
					delete(windowCounts, c)
				}
				l++
			}
			r++
		}
	}

	return max
}
