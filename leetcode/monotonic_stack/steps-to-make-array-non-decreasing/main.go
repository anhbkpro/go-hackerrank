package steps_to_make_array_non_decreasing

import (
	"math"
)

type stack []int

func (s *stack) pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func TotalSteps(nums []int) int {
	n := len(nums)
	ans := 0
	st := stack{}
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) != 0 && nums[i] > nums[st.top()] {
			dp[i] = int(math.Max(float64(dp[i]+1), float64(dp[st.pop()])))
			ans = int(math.Max(float64(ans), float64(dp[i])))
		}
		st.push(i)
	}

	return ans
}
