package valid_parentheses

func IsValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	r := []rune(s)
	var stack []string
	closeToOpenMap := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	for _, v := range r {
		c := string(v)
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack, c)
		} else {
			if stack == nil || len(stack) < 1 ||
				stack[len(stack)-1] != closeToOpenMap[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}

	return true
}
