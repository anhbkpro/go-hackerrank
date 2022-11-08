package find_all_anagrams_in_a_string

// FindAnagrams Same solution to leetcode/sliding-window/permutation-in-string/main.go
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var ans []int
	walker := 0
	freqP := freqOf(p)
	freqW := make(map[rune]int)
	runes := []rune(s)
	for runner, rr := range runes {
		freqW[rr]++
		if runner-walker+1 == len(p) && areMapsEqual(freqP, freqW) {
			ans = append(ans, walker)
		}

		if runner-walker+1 >= len(p) {
			lr := runes[walker]
			freqW[lr]--
			if freqW[lr] == 0 {
				delete(freqW, lr)
			}
			walker++
		}
	}
	return ans
}

func freqOf(s string) map[rune]int {
	res := make(map[rune]int)
	for _, r := range []rune(s) {
		res[r]++
	}
	return res
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
