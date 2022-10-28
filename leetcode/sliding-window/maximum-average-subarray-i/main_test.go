package maximum_average_subarray_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	assert.Equal(t, -1.00000, FindMaxAverage([]int{-1}, 1))
	assert.Equal(t, 12.75000, FindMaxAverageV2([]int{1, 12, -5, -6, 50, 3}, 4))
	assert.Equal(t, 5.00000, FindMaxAverage([]int{5}, 1))
}
