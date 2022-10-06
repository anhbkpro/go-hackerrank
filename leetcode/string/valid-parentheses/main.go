package valid_parentheses

func IsValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	r := []rune(s)
	var open []string
	m := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	for _, v := range r {
		if string(v) == "{" || string(v) == "(" || string(v) == "[" {
			open = append(open, string(v))
		} else {
			if open == nil || len(open) < 1 {
				return false
			}
			a := m[string(v)]
			e := open[len(open)-1]
			if a != e {
				return false
			}
			open = open[:len(open)-1]
		}
	}
	if len(open) > 0 {
		return false
	}

	return true
}
