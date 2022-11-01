package next_greater_element_ii

type stack []int

func (s *stack) Push(item int) {
	*s = append(*s, item)
}

func (s *stack) Pop() int {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) Peek() int {
	return (*s)[len(*s)-1]
}

func NextGreaterElements(nums []int) []int {
	s := stack{}
	n := len(nums)
	res := make([]int, n)
	for i := 2*len(nums) - 1; i >= 0; i-- {
		ans := -1
		for len(s) != 0 && nums[s.Peek()] <= nums[i%n] {
			s.Pop()
		}

		if len(s) != 0 {
			ans = nums[s.Peek()]
		}
		res[i%n] = ans
		s.Push(i % n)
	}
	return res
}
