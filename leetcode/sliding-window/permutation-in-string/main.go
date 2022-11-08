package permutation_in_string

// CheckInclusion Same solution to leetcode/sliding-window/find-all-anagrams-in-a-string/main.go
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) == len(s2) && len(s2) == 0 {
		return true
	}
	if len(s2) < len(s1) {
		return false
	}
	freqS1 := mapOf(s1)
	l, r := 0, 0
	freqWindow := make(map[string]int)
	for r < len(s2) {
		cRight := string(s2[r])
		freqWindow[cRight]++
		if r-l+1 == len(s1) {
			if areMapsEqual(freqS1, freqWindow) {
				return true
			}
		}

		if r-l+1 < len(s1) {
			r++
		} else {
			cLeft := string(s2[l])
			freqWindow[cLeft]--
			if freqWindow[cLeft] == 0 {
				delete(freqWindow, cLeft)
			}
			l++
			r++
		}
	}
	return false
}

func mapOf(str string) map[string]int {
	m := make(map[string]int)
	for _, s := range str {
		m[string(s)]++
	}
	return m
}

func areMapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
