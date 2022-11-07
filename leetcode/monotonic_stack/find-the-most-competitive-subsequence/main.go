package find_the_most_competitive_subsequence

import (
	"fmt"
)

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
	return len(*s) == 0
}

func MostCompetitive(nums []int, k int) []int {
	st := stack{}
	ans := make([]int, k)
	for i := 0; i < len(nums); i++ {
		for !st.empty() && nums[i] < nums[st.top()] && len(st)-1+len(nums)-i >= k {
			st.pop()
		}
		st.push(i)
	}
	fmt.Println(st)
	for len(st) > k {
		st.pop()
	}
	for i := len(ans) - 1; i >= 0; i-- {
		if !st.empty() {
			ans[i] = nums[st.pop()]
		}
	}
	return ans
}
