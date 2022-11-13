package squares_of_a_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	assert.Equal(t, []int{0, 1, 9, 16, 100}, SortedSquares([]int{-4, -1, 0, 3, 10}))
}
