package longest_repeating_character_replacement

import "math"

func CharacterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) < k {
		return len(s)
	}

	l, ans, mostFreq := 0, 0, 0
	freqWindow := make(map[string]int)
	for r := 0; r < len(s); r++ {
		cRight := string(s[r])
		freqWindow[cRight]++
		mostFreq = int(math.Max(float64(mostFreq), float64(freqWindow[cRight])))

		if r-l+1-mostFreq > k {
			cLeft := string(s[l])
			freqWindow[cLeft]--
			if freqWindow[cLeft] == 0 {
				delete(freqWindow, cLeft)
			}
			l++
		}
		ans = int(math.Max(float64(ans), float64(r-l+1)))
	}
	return ans
}
