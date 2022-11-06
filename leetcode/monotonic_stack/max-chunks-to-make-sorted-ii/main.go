package max_chunks_to_make_sorted_ii

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

func (s *stack) empty() bool {
	if len(*s) != 0 {
		return false
	}

	return true
}

// MaxChunksToSorted Works for both 769. Max Chunks To Make Sorted and 768. Max Chunks To Make Sorted II.
func MaxChunksToSorted(arr []int) int {
	st := stack{}
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		largest := num
		for !st.empty() && num < st.top() {
			largest = int(math.Max(float64(largest), float64(st.pop())))
		}
		st.push(largest)
	}
	return len(st)
}
