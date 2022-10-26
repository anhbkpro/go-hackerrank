package longest_repeating_character_replacement

import "math"

func CharacterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) < k {
		return len(s)
	}

	l, r, longestStr, maxFreq := 0, 0, 0, 0
	freqWindow := make(map[string]int)

	for r < len(s) {
		cRight := string(s[r])
		if _, ok := freqWindow[cRight]; ok {
			freqWindow[cRight]++
		} else {
			freqWindow[cRight] = 1
		}
		maxFreq = int(math.Max(float64(maxFreq), float64(freqWindow[cRight])))
		if r-l+1-maxFreq <= k {
			longestStr = int(math.Max(float64(longestStr), float64(r-l+1)))
			r++
		} else {
			cLeft := string(s[l])
			freqWindow[cLeft]--
			if freqWindow[cLeft] == 0 {
				delete(freqWindow, cLeft)
			}
			l++
			r++
		}
	}
	return longestStr
}
