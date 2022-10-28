package two_sum_ii_input_array_is_sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{1, 2}, TwoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 3}, TwoSum([]int{2, 3, 4}, 6))
	assert.Equal(t, []int{1, 2}, TwoSum([]int{-1, 0}, -1))
}
