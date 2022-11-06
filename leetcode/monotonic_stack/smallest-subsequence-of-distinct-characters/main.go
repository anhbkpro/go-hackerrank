package smallest_subsequence_of_distinct_characters

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

func (s *stack) empty() bool {
	if len(*s) != 0 {
		return false
	}

	return true
}

// SmallestSubsequence Works for both 1081. Smallest Subsequence of Distinct Characters
// and 316. Remove Duplicate Letters.
func SmallestSubsequence(s string) string {
	ans := ""
	runes := []rune(s)
	freq := make(map[rune]int)
	st := stack{}
	visited := make(map[rune]bool)
	for _, v := range runes {
		freq[v]++
	}
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		freq[char]--
		if visited[char] {
			continue
		}
		for !st.empty() && char < runes[st.top()] && freq[runes[st.top()]] > 0 {
			visited[runes[st.pop()]] = false
		}
		st.push(i)
		visited[runes[i]] = true
	}

	for !st.empty() {
		ans = string(runes[st.pop()]) + ans
	}
	return ans
}
