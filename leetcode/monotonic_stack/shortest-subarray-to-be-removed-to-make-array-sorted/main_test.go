package shortest_subarray_to_be_removed_to_make_array_sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLengthOfShortestSubarray(t *testing.T) {
	assert.Equal(t, 3, FindLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
}
