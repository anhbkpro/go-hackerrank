package shortest_unsorted_continuous_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	assert.Equal(t, 4, FindUnsortedSubarray([]int{1, 3, 2, 2, 2}))
	assert.Equal(t, 5, FindUnsortedSubarray([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, 2, FindUnsortedSubarray([]int{1, 3, 2, 4, 5}))
	assert.Equal(t, 2, FindUnsortedSubarray([]int{2, 1}))
	assert.Equal(t, 5, FindUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	assert.Equal(t, 0, FindUnsortedSubarray([]int{1, 2, 3, 4}))
	assert.Equal(t, 0, FindUnsortedSubarray([]int{1}))
}
