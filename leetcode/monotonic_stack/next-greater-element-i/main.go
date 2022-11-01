package next_greater_element_i

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

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	s := stack{} // decreasing monotonic
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		for len(s) != 0 && s.Peek() < num {
			m[s.Pop()] = num // pop until find the larger element in the stack or stack empty ==> make sure decreasing monotonic
		}
		s.Push(num)
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := m[nums1[i]]; ok {
			nums1[i] = v
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
