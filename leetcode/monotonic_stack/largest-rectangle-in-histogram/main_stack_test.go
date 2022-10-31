package largest_rectangle_in_histogram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleAreaV2(t *testing.T) {
	//assert.Equal(t, 4, LargestRectangleAreaByStack([]int{2, 4}))
	assert.Equal(t, 10, LargestRectangleAreaByStack([]int{2, 1, 5, 6, 2, 3}))
}
