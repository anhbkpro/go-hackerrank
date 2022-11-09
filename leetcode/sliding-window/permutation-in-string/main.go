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
	walker := 0
	freqW := make(map[rune]int)
	runes := []rune(s2)
	for runner, rr := range runes {
		freqW[rr]++
		if runner-walker+1 == len(s1) && areMapsEqual(freqS1, freqW) {
			return true
		}

		if runner-walker+1 >= len(s1) {
			lr := runes[walker]
			freqW[lr]--
			if freqW[lr] == 0 {
				delete(freqW, lr)
			}
			walker++
		}
	}
	return false
}

func mapOf(str string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range []rune(str) {
		m[r]++
	}
	return m
}

func areMapsEqual(m1, m2 map[rune]int) bool {
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
