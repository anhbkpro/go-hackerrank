package find_all_anagrams_in_a_string

// FindAnagrams Same solution to leetcode/sliding-window/permutation-in-string/main.go
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var ans []int
	l, r := 0, 0
	freqP := freqOf(p)
	freqWindow := make(map[string]int)
	for r < len(s) {
		cRight := string(s[r])
		freqWindow[cRight]++
		if r-l+1 == len(p) && areMapsEqual(freqP, freqWindow) {
			ans = append(ans, l)
		}

		if r-l+1 < len(p) {
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
	return ans
}

func freqOf(s string) map[string]int {
	res := make(map[string]int)
	for _, item := range s {
		res[string(item)]++
	}
	return res
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
