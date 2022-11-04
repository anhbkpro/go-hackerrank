package remove_one_element_to_make_the_array_strictly_increasing

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

func CanBeIncreasing(nums []int) bool {
	if len(nums) < 3 {
		return true
	}

	st := stack{}
	removedIndex := -1 // decide which item will be removed: item on top of stack or current item
	for i := 0; i < len(nums); i++ {
		if removedIndex >= 0 {
			break
		}
		for len(st) != 0 && nums[i] <= nums[st.top()] {
			removedIndex = st.top()
			st.pop()
			if len(st) != 0 && nums[i] <= nums[st.top()] { // if more than 1 items on the stack larger than current item, remove current item
				removedIndex = i
			}
			break
		}
		st.push(i)
	}

	for i := 0; i < len(nums)-1; i++ {
		if i >= removedIndex {
			nums[i] = nums[i+1]
		} else {
			nums[i] = nums[i]
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}
