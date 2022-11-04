package largest_rectangle_in_histogram

type indexElement struct {
	index  int
	height int
}
type stack []indexElement

func (s *stack) push(element indexElement) {
	*s = append(*s, element)
}

func (s *stack) pop() indexElement {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) top() indexElement {
	return (*s)[len(*s)-1]
}

func LargestRectangleAreaByStack(heights []int) int {
	start := 0
	st := stack{}
	max := 0

	for i, h := range heights {
		start = i
		for len(st) != 0 && st.top().height > h {
			item := st.pop()
			area := (i - item.index) * item.height
			if area > max {
				max = area
			}
			start = item.index
		}
		st.push(indexElement{
			index:  start,
			height: h,
		})
	}

	for len(st) != 0 {
		item := st.pop()
		area := (len(heights) - item.index) * item.height
		if area > max {
			max = area
		}
	}

	return max
}
