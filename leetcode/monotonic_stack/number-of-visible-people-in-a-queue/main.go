package number_of_visible_people_in_a_queue

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

// CanSeePersonsCount
// My solution is using decreasing monotonic.
// The key thing is how to find the nearest higher people to the right.
// If not found, returns 0.
// If found, the number of people can see (to the right) = index of top of stack (nearest higher people to the right) - current index.
func CanSeePersonsCount(heights []int) []int {
	st := stack{}
	n := len(heights)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for !st.empty() && heights[i] > heights[st.top()] {
			ans[i]++
			st.pop()
		}
		if !st.empty() {
			ans[i]++
		}
		st.push(i)
	}
	return ans
}
