package remove_one_element_to_make_the_array_strictly_increasing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanBeIncreasing(t *testing.T) {
	assert.Equal(t, true, CanBeIncreasing([]int{105, 924, 32, 968}))
	assert.Equal(t, true, CanBeIncreasing([]int{100, 21, 100}))
	assert.Equal(t, true, CanBeIncreasing([]int{1, 2, 10, 5, 7}))
	assert.Equal(t, false, CanBeIncreasing([]int{2, 3, 1, 2}))
	assert.Equal(t, false, CanBeIncreasing([]int{1, 1, 1}))
}
