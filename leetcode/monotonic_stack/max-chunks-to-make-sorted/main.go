package max_chunks_to_make_sorted

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

func MaxChunksToSortedCleverSolution(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	count := 0
	max := 0
	for i := 0; i < n; i++ {
		max = int(math.Max(float64(max), float64(arr[i])))
		if max == i {
			count++
		}
	}
	return count
}
