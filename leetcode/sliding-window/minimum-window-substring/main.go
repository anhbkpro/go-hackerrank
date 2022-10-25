package minimum_window_substring

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	// Map which keeps a count of all the unique characters in t.
	mapT := mapOf(t)

	// Number of unique characters in t, which need to be present in the desired window.
	required := len(mapT)
	formed := 0
	r, l := 0, 0
	windowCounts := make(map[string]int)
	ans := []int{-1, 0, 0}
	for r < len(s) {
		// Add one character from the right to the window
		c := string(s[r])
		if _, ok := windowCounts[c]; ok {
			windowCounts[c]++
		} else {
			windowCounts[c] = 1
		}
		if v1, v2 := windowCounts[c], mapT[c]; v1 != 0 && v2 != 0 && v1 == v2 {
			formed++
		}
		for l <= r && formed == required {
			c = string(s[l])
			// Save the smallest window until now.
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}
			windowCounts[c]--
			if v1, v2 := windowCounts[c], mapT[c]; v1 < v2 {
				formed--
			}
			l++
		}
		r++
	}

	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
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
