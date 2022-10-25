package permutation_in_string

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) == len(s2) && len(s2) == 0 {
		return true
	}
	if len(s2) < len(s1) {
		return false
	}
	freqS1 := mapOf(s1)
	l, r := 0, 0
	freqS2 := make(map[string]int)
	for r < len(s2) {
		cRight := string(s2[r])
		if _, ok := freqS2[cRight]; ok {
			freqS2[cRight]++
		} else {
			freqS2[cRight] = 1
		}

		if r-l+1 == len(s1) {
			if areMapsEqual(freqS1, freqS2) {
				return true
			}
		}

		if r-l+1 < len(s1) {
			r++
		} else {
			cLeft := string(s2[l])
			freqS2[cLeft]--
			if freqS2[cLeft] == 0 {
				delete(freqS2, cLeft)
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
		if _, ok := m[string(s)]; ok {
			m[string(s)]++
		} else {
			m[string(s)] = 1
		}
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
