package main

/*
 * Complete the 'IsBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func IsBalanced(s string) string {
	var stack Stack
	for _, r := range s {
		if isOpenBrace(r) {
			stack.Push(string(r)) // a = string(65)
		} else if isClosedBrace(r) {
			item, ok := stack.Pop()
			if !ok {
				return "NO"
			}
			if isMissing(item, string(r)) {
				return "NO"
			}
		}
	}

	if len(stack) > 0 {
		return "NO"
	}

	return "YES"
}

type Stack []string

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isOpenBrace(c rune) bool {
	for _, v := range "([{" {
		if v == c {
			return true
		}
	}

	return false
}

func isClosedBrace(c rune) bool {
	for _, v := range ")]}" {
		if v == c {
			return true
		}
	}

	return false
}

func isMissing(open, close string) bool {
	openToCloseMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	if v, ok := openToCloseMap[open]; ok {
		return v != close
	}

	return false
}
