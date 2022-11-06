package next_greater_element_i

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

// NextGreaterElement What makes this problem EASY is nums1 and nums2 are two distinct 0-indexed integer arrays
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	s := stack{} // decreasing monotonic
	nextGreaterElementMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		for len(s) != 0 && s.top() < num {
			// because num is distinct ==> so that the next greater element of each element can not be overridden
			nextGreaterElementMap[s.pop()] = num // pop until find an element in the stack is less than the current one or stack empty ==> make sure decreasing monotonic
		}
		s.push(num)
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := nextGreaterElementMap[nums1[i]]; ok {
			nums1[i] = v
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
