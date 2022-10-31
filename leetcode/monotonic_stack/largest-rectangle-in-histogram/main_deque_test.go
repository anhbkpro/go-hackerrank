package largest_rectangle_in_histogram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	assert.Equal(t, 4, LargestRectangleArea([]int{2, 4}))
	assert.Equal(t, 10, LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
