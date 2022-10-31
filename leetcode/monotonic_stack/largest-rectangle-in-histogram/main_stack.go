package largest_rectangle_in_histogram

type indexElement struct {
	index  int
	height int
}
type stack []indexElement

func (s *stack) Push(element indexElement) {
	*s = append(*s, element)
}

func (s *stack) Pop() indexElement {
	returnVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return returnVal
}

func (s *stack) Peek() indexElement {
	return (*s)[len(*s)-1]
}

func LargestRectangleAreaByStack(heights []int) int {
	i := 0
	stack := stack{}
	max := 0

	for i < len(heights) {
		currentHeight := heights[i]
		if len(stack) == 0 || currentHeight > stack.Peek().height { // stack contains increasing index_height
			stack.Push(indexElement{
				index:  i,
				height: currentHeight,
			})
		} else {
			previous := i
			for len(stack) > 0 && currentHeight <= stack.Peek().height {
				currentIndex := i
				poppedVal := stack.Pop()
				localArea := (currentIndex - poppedVal.index) * poppedVal.height
				if localArea > max {
					max = localArea
				}
				previous = poppedVal.index
			}
			stack.Push(indexElement{
				index:  previous,
				height: heights[i],
			})
		}
		i++
	}

	for len(stack) > 0 {
		poppedVal := stack.Pop()
		currentWith := i - poppedVal.index
		area := currentWith * poppedVal.height
		if area > max {
			max = area
		}
	}
	return max
}
