package next_greater_element_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	assert.Equal(t, []int{-1, 5, 5, 5, 5}, NextGreaterElements([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, []int{2, -1, 2}, NextGreaterElements([]int{1, 2, 1}))
	//assert.Equal(t, []int{2, 3, 4, -1, 4}, NextGreaterElements([]int{1, 2, 3, 4, 3}))
}
