package _32_pattern

import "math"

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

func Find132pattern(nums []int) bool {
	st := stack{}

	prev := math.MinInt
	for i := len(nums) - 1; i >= 0; i-- {
		if prev > nums[i] {
			return true
		}

		for len(st) != 0 && st.top() < nums[i] {
			prev = st.top()
			st.pop()
		}

		st.push(nums[i])
	}

	return false
}
