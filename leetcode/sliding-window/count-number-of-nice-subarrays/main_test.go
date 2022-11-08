package count_number_of_nice_subarrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfSubarrays(t *testing.T) {
	assert.Equal(t, 2, NumberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	assert.Equal(t, 0, NumberOfSubarrays([]int{2, 4, 6}, 1))
	assert.Equal(t, 16, NumberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
