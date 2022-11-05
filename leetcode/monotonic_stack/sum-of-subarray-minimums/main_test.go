package sum_of_subarray_minimums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSubarrayMins(t *testing.T) {
	assert.Equal(t, 17, SumSubarrayMins([]int{3, 1, 2, 4}))
	assert.Equal(t, 444, SumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
