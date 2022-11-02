package shortest_unsorted_continuous_subarray

import (
	"math"
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

func (s *stack) clear() {
	*s = make([]int, 0)
}

func FindUnsortedSubarray(nums []int) int {
	st := stack{}
	l, r := len(nums), 0
	for i := 0; i < len(nums); i++ {
		for len(st) != 0 && nums[st.top()] > nums[i] {
			l = int(math.Min(float64(l), float64(st.pop())))
		}
		st.push(i)
	}
	st.clear()
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) != 0 && nums[i] > nums[st.top()] {
			r = int(math.Max(float64(r), float64(st.pop())))
		}
		st.push(i)
	}
	if r > l {
		return r - l + 1
	}
	return 0
}
