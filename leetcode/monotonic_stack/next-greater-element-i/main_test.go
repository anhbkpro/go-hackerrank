package next_greater_element_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	assert.Equal(t, []int{-1, 3, -1}, NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	assert.Equal(t, []int{3, -1}, NextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
