package max_subarrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSumSubArrayOfSizeK(t *testing.T) {
	input := []int32{2, 1, 5, 1, 3, 2}
	got := MaxSumSubArrayOfSizeK(3, input)
	assert.Equal(t, int32(9), got)

	input2 := []int32{2, 3, 4, 1, 5}
	got2 := MaxSumSubArrayOfSizeK(2, input2)
	assert.Equal(t, int32(7), got2)
}
