package longest_substring_without_repeating_characters

import "math"

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	l, r := 0, 0
	formed := true
	windowCounts := make(map[string]int)
	ans := []int{math.MaxInt, 0, 0}
	for r < len(s) {
		cr := string(s[r])
		if _, ok := windowCounts[cr]; ok {
			windowCounts[cr]++
			formed = false
		} else {
			windowCounts[cr] = 1
		}
		for (l <= r && !formed) || r == len(s)-1 {

			cl := string(s[l])
			if ans[0] == math.MaxInt || r-l > ans[0] {
				ans[0] = r - l
				ans[1] = l
				ans[2] = r
			}

			l++
			if cl == cr {
				windowCounts[cl]--
				formed = true
				break
			}
			delete(windowCounts, cl)
		}
		r++
	}
	if ans[0] == math.MaxInt {
		return len(s)
	}

	return ans[0]
}

func LengthOfLongestSubstringDiscussion(s string) int {
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
