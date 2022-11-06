package remove_duplicate_letters

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

func (s *stack) clear() {
	*s = make([]int, 0)
}

// RemoveDuplicateLetters Works for both 1081. Smallest Subsequence of Distinct Characters
// and 316. Remove Duplicate Letters.
func RemoveDuplicateLetters(s string) string {
	st := stack{}
	runes := []rune(s)
	freqChar := make(map[rune]int)
	visited := make(map[rune]bool)
	for _, r := range []rune(s) {
		freqChar[r]++
	}
	for i := 0; i < len(runes); i++ {
		freqChar[runes[i]]--
		if visited[runes[i]] {
			continue
		}
		// if there is a still good enough character (lexicographical order), remove top item of stack and push better one
		// the top of stack should be behind the current one // bcabc
		// draw diagram to better visualization
		for len(st) != 0 && runes[i] < runes[st.top()] && freqChar[runes[st.top()]] > 0 { // stack = |b|c, current = a
			visited[runes[st.top()]] = false // mark c and b are unvisited
			st.pop()                         // pop c and b out of stack
		}
		st.push(i) // push a
		visited[runes[i]] = true
	}
	res := ""
	for len(st) != 0 {
		res = string(runes[st.pop()]) + res
	}
	return res
}
