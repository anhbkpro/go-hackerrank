package trapping_rain_water

import (
	"fmt"
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

func Trap(height []int) int {
	st := stack{}
	ans := 0
	for current := 0; current < len(height); current++ {
		currentHeight := height[current]
		for len(st) != 0 && height[current] > height[st.top()] {
			top := st.top()
			st.pop()
			if len(st) == 0 {
				break
			}
			boundedHeight := int(math.Min(float64(currentHeight), float64(height[st.top()]))) - height[top]
			distance := current - st.top() - 1
			fmt.Printf("Trapped water between %d and %d is %d \n", st.top(), current, boundedHeight*distance)
			ans += boundedHeight * distance
		}
		// keep adding index to stack if bar is smaller than or equal to the bar at top of stack
		st.push(current)
	}

	return ans
}
