package number_of_sub_arrays_of_size_k_and_average_greater_than_or_equal_to_threshold

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumOfSubarrays(t *testing.T) {
	assert.Equal(t, 3, NumOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
	assert.Equal(t, 6, NumOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5))
}
