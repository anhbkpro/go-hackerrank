package sum_of_subarray_ranges

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

func SubArrayRanges(nums []int) int64 {
	n := len(nums)
	var answer int64
	stk := stack{}
	// Find the sum of all the minimum
	for right := 0; right <= n; right++ {
		for !stk.empty() && (right == n || nums[stk.top()] >= nums[right]) {
			mid := stk.top()
			stk.pop()
			left := 0
			if stk.empty() {
				left = -1
			} else {
				left = stk.top()
			}
			answer -= int64(nums[mid] * (right - mid) * (mid - left))
		}
		stk.push(right)
	}

	// Find the sum of all the maximum
	stk.pop()
	for right := 0; right <= n; right++ {
		for !stk.empty() && (right == n || nums[stk.top()] <= nums[right]) {
			mid := stk.top()
			stk.pop()
			left := 0
			if stk.empty() {
				left = -1
			} else {
				left = stk.top()
			}
			answer += int64(nums[mid] * (right - mid) * (mid - left))
		}
		stk.push(right)
	}
	return answer
}
