package subarray_sum_equals_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	got := SubarraySum([]int{1, 1, 1}, 2)
	assert.Equal(t, 2, got)

	got2 := SubarraySum([]int{1, 2, 3}, 3)
	assert.Equal(t, 2, got2)
}
