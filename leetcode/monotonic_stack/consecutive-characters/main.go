package consecutive_characters

type stack []int

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	if len(*s) != 0 {
		return false
	}

	return true
}

func MaxPower(s string) int {
	stk := stack{}
	ans := 1
	runes := []rune(s)
	for i := 0; i <= len(runes); i++ {
		for !stk.empty() && (i == len(runes) || runes[i] != runes[stk.top()]) {
			if len(stk) > ans {
				ans = len(stk)
			}
			stk.pop()
		}
		stk.push(i)
	}
	return ans
}
